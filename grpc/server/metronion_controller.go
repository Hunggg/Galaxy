package server

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	pb "github.com/metrogalaxy-io/metrogalaxy-api/grpc/proto/metronion/v1"
	"github.com/metrogalaxy-io/metrogalaxy-api/services"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/metronion"

	"github.com/metrogalaxy-io/metrogalaxy-api/env"
	"github.com/metrogalaxy-io/metrogalaxy-api/inject"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/cache"
	"github.com/vmihailenco/msgpack/v5"
	"go.uber.org/zap"
)

const (
	DefaultMetronionVersion              = uint64(0)
	DefaultCacheTime                     = 20 * time.Second
	ENV_METROGALAXY_API_PUBLISH_METADATA = "METROGALAXY_API_PUBLISH_METADATA"
)

var (
	ListStats = []string{"creativity", "charisma", "resolve", "fitness", "intellect"}
)

type MetronionController struct {
	l *zap.SugaredLogger
	// pb.UnimplementedMetronionServer
	cache cache.Cache
	db    metronion.PersistentDb
	// metadataStorage *metadata.MetadataStorage
	config env.Config
}

func NewMetronionController(injector *inject.Injector) *MetronionController {
	// metadataStorage := injector.ProvideMetadataStorage()
	cache := injector.ProvideCache()
	db := injector.ProvideMetronionDb()
	config := env.GetConfig()

	return &MetronionController{
		l:     zap.S(),
		cache: cache,
		db:    db,
		// metadataStorage: metadataStorage,
		config: config,
	}
}

func (c *MetronionController) GetMetadata(ctx context.Context, req *pb.GetMetadataRequest) (*pb.GetMetadataResponse, error) {
	var res = &pb.GetMetadataResponse{}
	if err := req.ValidateAll(); err != nil {
		c.l.Debugw("GetMetadata invalid request", "error", err)
		return nil, err
	}
	metronionID := req.GetId().GetValue()

	// get from cache first
	cacheKey := fmt.Sprintf("metronion_metadata_%v", metronionID)
	cacheRes, ok := c.cache.Get(cacheKey)
	if ok {
		if err := msgpack.Unmarshal([]byte(cacheRes.(string)), res); err != nil {
			c.l.Errorw("error unmarshal metronion cache response", "error", err)
			return nil, err
		}
		return res, nil
	}

	metronion, err := c.db.GetMetronionById(metronionID)
	if err != nil {
		c.l.Errorw("error get metronion metadata from db", "error", err)
		return nil, err
	}
	res = c.convertMetronionMetadata(metronion)

	bRes, err := msgpack.Marshal(metronion)
	if err != nil {
		c.l.Errorw("error marshal metronion cache response", "error", err)
		return nil, err
	}
	// save to cache
	_ = c.cache.Set(cacheKey, string(bRes), DefaultCacheTime)
	return res, nil
}

func (c *MetronionController) GetListMetadata(ctx context.Context, req *pb.GetListMetadataRequest) (*pb.GetListMetadataResponse, error) {
	var res = &pb.GetListMetadataResponse{}
	const (
		DefaultLimit = 6
		DefaultSort  = "lowest_id"
	)

	if err := req.ValidateAll(); err != nil {
		c.l.Debugw("GetListMetadata invalid request", "error", err)
		return nil, err
	}

	account := req.GetAccount().GetValue()
	sortType := req.GetSort()
	offset := req.GetOffset()
	limit := req.GetLimit()
	status := req.GetStatus()
	gender := req.GetGender()

	if limit == 0 {
		limit = DefaultLimit
	}
	if sortType == "" {
		sortType = DefaultSort
	}

	var stat = make(map[string][]uint64)
	for key, value := range req.GetStat() {
		if err := isValidStatsName(key); err != nil {
			c.l.Debugw("GetListMetadata invalid request", "error", err)
			return nil, err
		}
		stat[key] = make([]uint64, 2)
		values := value.GetValues()
		if len(values) != 2 {
			err := fmt.Errorf(fmt.Sprintf("invalid stat `%s` request", key))
			c.l.Debugw("GetListMetadata invalid request", "error", err)
			return nil, err
		}
		for i, item := range value.GetValues() {
			stat[key][i] = uint64(item.GetNumberValue())
		}
	}
	var metronionId *uint64
	if req.GetId() != nil {
		id := req.GetId().GetValue()
		metronionId = &id
	}

	filter := metronion.MetronionFilterParams{
		Owner:       account,
		Sort:        sortType,
		MetronionID: metronionId,
		Status:      status,
		Gender:      gender,
		Stat:        stat,
	}
	cacheKey := composeGetListMetadataCacheKey(int(offset), int(limit), filter)
	cacheRes, ok := c.cache.Get(cacheKey)
	if ok {
		if err := msgpack.Unmarshal([]byte(cacheRes.(string)), res); err != nil {
			c.l.Errorw("error unmarshal metronion cache response", "error", err)
			return nil, err
		}
		return res, nil
	}

	listMetronion, count, err := c.db.GetListMetronions(int(offset), int(limit), filter)
	if err != nil {
		c.l.Errorw("error get list metronions from database", "error", err)
		return nil, err
	}

	listConvertedMetronion, err := c.convertMetronionMetadataArray(listMetronion)
	if err != nil {
		c.l.Errorw("error convert list metronions", "error", err)
		return nil, err
	}

	res.Count = count
	res.Data = listConvertedMetronion

	// save to cache
	bRes, err := msgpack.Marshal(res)
	if err != nil {
		c.l.Errorw("error marshal metronion cache response", "error", err)
		return nil, err
	}
	_ = c.cache.Set(cacheKey, string(bRes), DefaultCacheTime)
	return res, nil
}

func (c *MetronionController) GetMetronionActivities(ctx context.Context, req *pb.GetMetronionActivitiesRequest) (*pb.GetMetronionActivitiesResponse, error) {
	var res = &pb.GetMetronionActivitiesResponse{}
	if err := req.ValidateAll(); err != nil {
		c.l.Debugw("GetMetronionActivities invalid request", "error", err)
		return nil, err
	}
	metronionID := req.GetId().GetValue()
	sortType := req.GetSort()
	if sortType == "" {
		sortType = metronion.SortTypeNewest
	}

	// get from cache first
	cacheKey := fmt.Sprintf("metronion_activities_%v_%v", metronionID, sortType)
	cacheRes, ok := c.cache.Get(cacheKey)
	if ok {
		if err := msgpack.Unmarshal([]byte(cacheRes.(string)), res); err != nil {
			c.l.Errorw("error unmarshal metronion activities response", "error", err)
			return nil, err
		}
		return res, nil
	}

	activities, err := c.db.GetMetronionActivities(metronionID, sortType)
	if err != nil {
		c.l.Errorw("error get metronion activities from db", "error", err)
		return nil, err
	}
	res.Data = c.convertMetronionActivities(activities)

	bRes, err := msgpack.Marshal(res)
	if err != nil {
		c.l.Errorw("error marshal metronion activities cache response", "error", err)
		return nil, err
	}
	// save to cache
	_ = c.cache.Set(cacheKey, string(bRes), DefaultCacheTime)
	return res, nil
}

func (c *MetronionController) GetMetronionListing(ctx context.Context, req *pb.GetMetronionListingRequest) (*pb.GetMetronionListingResponse, error) {
	var res = &pb.GetMetronionListingResponse{}
	if err := req.ValidateAll(); err != nil {
		c.l.Debugw("GetMetronionListing invalid request", "error", err)
		return nil, err
	}
	metronionID := req.GetId().GetValue()

	// get from cache first
	cacheKey := fmt.Sprintf("metronion_listing_%v", metronionID)
	cacheRes, ok := c.cache.Get(cacheKey)
	if ok {
		if err := msgpack.Unmarshal([]byte(cacheRes.(string)), res); err != nil {
			c.l.Errorw("error unmarshal metronion listing response", "error", err)
			return nil, err
		}
		return res, nil
	}

	listing, err := c.db.GetMetronionListing(metronionID)
	if err != nil {
		c.l.Errorw("error get metronion listing from db", "error", err)
		return nil, err
	}
	res.Data = c.convertMetronionListing(listing)
	bRes, err := msgpack.Marshal(res)
	if err != nil {
		c.l.Errorw("error marshal metronion listing cache response", "error", err)
		return nil, err
	}
	// save to cache
	_ = c.cache.Set(cacheKey, string(bRes), DefaultCacheTime)
	return res, nil
}

func (c *MetronionController) GetMetronionOffers(ctx context.Context, req *pb.GetMetronionOffersRequest) (*pb.GetMetronionOffersResponse, error) {
	var res = &pb.GetMetronionOffersResponse{}
	if err := req.ValidateAll(); err != nil {
		c.l.Debugw("GetMetronionOffers invalid request", "error", err)
		return nil, err
	}
	metronionID := req.GetId().GetValue()
	sortType := req.GetSort()

	// get from cache first
	cacheKey := fmt.Sprintf("metronion_offers_%v_%v", metronionID, sortType)
	cacheRes, ok := c.cache.Get(cacheKey)
	if ok {
		if err := msgpack.Unmarshal([]byte(cacheRes.(string)), res); err != nil {
			c.l.Errorw("error unmarshal metronion offers response", "error", err)
			return nil, err
		}
		return res, nil
	}

	offers, err := c.db.GetMetronionOffers(metronionID, sortType)
	if err != nil {
		c.l.Errorw("error get metronion offers from db", "error", err)
		return nil, err
	}
	res.Data = c.convertMetronionOffers(offers)

	bRes, err := msgpack.Marshal(res)
	if err != nil {
		c.l.Errorw("error marshal metronion cache response", "error", err)
		return nil, err
	}
	// save to cache
	_ = c.cache.Set(cacheKey, string(bRes), DefaultCacheTime)
	return res, nil
}

func composeGetListMetadataCacheKey(offset, limit int, filterParams metronion.MetronionFilterParams) string {
	// cache key format: metronion_list_metadata_<offset>_<limit>_<account>_<sort>_<metronion_id>_<status>_<gender>_<stats>
	var sb strings.Builder
	sb.WriteString("metronion_list_metadata")
	sb.WriteString(fmt.Sprintf("_%v", offset))
	sb.WriteString(fmt.Sprintf("_%v", limit))
	sb.WriteString(fmt.Sprintf("_%v", filterParams.Owner))
	sb.WriteString(fmt.Sprintf("_%v", filterParams.Sort))
	if filterParams.MetronionID != nil {
		sb.WriteString(fmt.Sprintf("_%v", *filterParams.MetronionID))
	} else {
		sb.WriteString("_nil")
	}
	sb.WriteString(fmt.Sprintf("_%v", filterParams.Status))
	sb.WriteString(fmt.Sprintf("_%v", filterParams.Gender))

	for _, stat := range ListStats {
		statValueString := make([]string, len(filterParams.Stat[stat]))
		for i := range filterParams.Stat[stat] {
			statValueString[i] = strconv.FormatUint(filterParams.Stat[stat][i], 10)
		}
		value := strings.Join(statValueString, "")
		sb.WriteString(fmt.Sprintf("_%v%v", stat, value))
	}

	return sb.String()
}

func (c *MetronionController) convertMetronionActivities(inputData []metronion.MetronionActivity) []*pb.MetronionActivity {
	var result = make([]*pb.MetronionActivity, len(inputData))
	for i := range inputData {
		result[i] = &pb.MetronionActivity{
			Id:          inputData[i].MetronionID,
			Type:        inputData[i].ActivityType,
			FromAccount: inputData[i].FromAccount,
			ToAccount:   inputData[i].ToAccount,
			Price:       inputData[i].Price,
			Timestamp:   inputData[i].Timestamp,
			BlockNumber: inputData[i].BlockNumber,
			TxHash:      inputData[i].TxHash,
		}
	}
	return result
}

func (c *MetronionController) convertMetronionListing(inputData metronion.MetronionListing) *pb.MetronionListing {
	return &pb.MetronionListing{
		Id:          inputData.MetronionID,
		FromAccount: inputData.FromAccount,
		Price:       inputData.Price,
		Timestamp:   inputData.Timestamp,
		BlockNumber: inputData.BlockNumber,
		TxHash:      inputData.TxHash,
	}
}

func (c *MetronionController) convertMetronionOffers(inputData []metronion.MetronionOffers) []*pb.MetronionOffers {
	var result = make([]*pb.MetronionOffers, len(inputData))
	for i := range inputData {
		result[i] = &pb.MetronionOffers{
			Id:          inputData[i].MetronionID,
			FromAccount: inputData[i].FromAccount,
			Price:       inputData[i].Price,
			Timestamp:   inputData[i].Timestamp,
			BlockNumber: inputData[i].BlockNumber,
			TxHash:      inputData[i].TxHash,
		}
	}
	return result
}

func (c *MetronionController) convertMetronionMetadataArray(inputData []metronion.Metronion) ([]*pb.GetMetadataResponse, error) {
	var result = make([]*pb.GetMetadataResponse, len(inputData))
	for i := range inputData {
		result[i] = c.convertMetronionMetadata(inputData[i])
	}
	return result, nil
}

func (c *MetronionController) convertMetronionMetadata(inputData metronion.Metronion) *pb.GetMetadataResponse {
	baseUri := strings.TrimSuffix(c.config.MetronionNFTConfig.ImageURI, "/")
	metronionServerId, err := c.getMetronionImageID(inputData.Id)
	var imageUri string
	if err != nil {
		imageUri = fmt.Sprintf("%s/unknown_metronion.png", baseUri)
	} else {
		imageUri = fmt.Sprintf("%s/%d.png", baseUri, metronionServerId)
	}

	var wearables = make([]*pb.Wearables, len(inputData.Wearables))
	for i, item := range inputData.Wearables {
		wearables[i] = &pb.Wearables{
			Id:         item.WearableId,
			Type:       item.Type,
			Gender:     item.Gender,
			Name:       item.Name,
			Changeable: item.Changeable,
			IsOrigin:   item.IsOrigin,
			IsRequired: item.IsRequired,
			Rarity:     item.Rarity,
			Image:      item.Image,
			Stats: &pb.MetronionStats{
				Creativity: item.Stats.Creativity,
				Charisma:   item.Stats.Charisma,
				Resolve:    item.Stats.Resolve,
				Fitness:    item.Stats.Fitness,
				Intellect:  item.Stats.Intellect,
			},
		}
	}

	return &pb.GetMetadataResponse{
		Id:                 inputData.Id,
		Name:               inputData.Name,
		VersionId:          inputData.VersionId,
		CreatedAtTimestamp: inputData.CreatedAtTimestamp,
		UpdatedAtTimestamp: inputData.UpdatedAtTimestamp,
		CreatedAtBlock:     inputData.CreatedAtBlock,
		UpdatedAtBlock:     inputData.UpdatedAtBlock,
		Owner:              inputData.Owner,
		LastPrice:          inputData.LastPrice,
		CurrentPrice:       inputData.CurrentPrice,
		TopBid:             inputData.TopBid,
		ListedBy:           inputData.ListedBy,
		Image:              imageUri,
		Gender:             inputData.Gender,
		Type:               inputData.Type,
		Star:               uint64(inputData.Star),
		SpecialStar:        uint64(inputData.SpecialStar),
		BaseStats:          c.convertStats(inputData.BaseStats),
		SpecialStats:       c.convertStats(inputData.SpecialStats),
		Wearables:          wearables,
	}
}

// func (c *MetronionController) getMetronionMetadata(id uint64) (nosql.Metronion, error) {
// 	metadata := nosql.Metronion{}
// 	var err error
// 	if viper.GetBool(ENV_METROGALAXY_API_PUBLISH_METADATA) {
// 		metadata, err = c.metadataStorage.GetMetronionById(id)
// 		if err != nil {
// 			return nosql.Metronion{}, err
// 		}
// 	}
// 	return metadata, nil
// }

// getMetronionImageID Mapping from MetronionID in contract with MetronionID in server
// formula: metronion_server_id = (metronion_nft_id + starting_index) % max_supply
func (c *MetronionController) getMetronionImageID(metronionNftID uint64) (uint64, error) {
	key := fmt.Sprintf("metronion_version_config_%d", DefaultMetronionVersion)
	versionConfig, ok := c.cache.Get(key)
	if !ok {
		config, err := c.db.GetMetronionVersionConfig(DefaultMetronionVersion)
		if err != nil {
			c.l.Errorw("error get metronion version config from database", "error", err)
			return 0, err
		}
		versionConfig = config
		if err := c.cache.Set(key, config, 0); err != nil {
			c.l.Errorw("error set cache version config", "error", err)
			return 0, err
		}
	}

	config := versionConfig.(metronion.MetronionVersionConfig)
	if config.StartingIndex == 0 {
		return 0, fmt.Errorf("starting index is not finalized yet")
	}

	return (metronionNftID + config.StartingIndex) % config.MaxSupply, nil
}

func (c *MetronionController) convertStats(stats services.Stats) *pb.MetronionStats {
	return &pb.MetronionStats{
		Creativity: stats.Creativity,
		Charisma:   stats.Charisma,
		Resolve:    stats.Resolve,
		Fitness:    stats.Fitness,
		Intellect:  stats.Intellect,
	}
}

func isValidStatsName(input string) error {
	var found = false
	for _, item := range ListStats {
		if input == item {
			found = true
		}
	}
	if !found {
		return fmt.Errorf("invalid stats name")
	}
	return nil
}