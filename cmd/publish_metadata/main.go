package publishmetadata

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
	cockroachdblib "github.com/metrogalaxy-io/metrogalaxy-api/libs/cockroachdb"
	"github.com/metrogalaxy-io/metrogalaxy-api/services"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/metronion"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/metronion/cockroachdb"
	"go.uber.org/zap"
)

const (
	WEARABLE_S3_URI = "s3://metrogalaxy-storage/metronion_wearables"
)

type PublishMetadataService struct {
	l            *zap.SugaredLogger
	persistentDb metronion.PersistentDb
}

func NewPublishMetadataService() *PublishMetadataService {
	l := zap.S()
	db, err := cockroachdblib.NewCockroachDbConnection()
	if err != nil {
		l.Panic(err)
	}
	client, err := cockroachdb.NewCockroachDB(db)
	if err != nil {
		l.Panic(err)
	}

	return &PublishMetadataService{
		l:            zap.S(),
		persistentDb: client,
	}
}

func (s *PublishMetadataService) Run(csvFile string) error {
	s.l.Infof("preparing to parse metronion metadata from CSV file %v", csvFile)
	metronions, err := s.ParseMetronionMetadata(csvFile)
	if err != nil {
		s.l.Errorw("error parse metronion metadata")
		return err
	}
	s.l.Infow("parse metronion metadata successfully", "metronions_length", len(metronions))

	convertedData, err := convertMetronion(metronions)
	if err != nil {
		s.l.Errorw("error convert metronion", "error", err)
		return err
	}
	s.l.Infow("convert metronion metadata successfully", "metronions_length", len(convertedData))
	if err := s.persistentDb.SaveBatchMetronionMetadata(convertedData); err != nil {
		s.l.Errorw("error save batch metronion to persistent database", "error", err)
		return err
	}
	s.l.Infof("save batch metronion successfully")

	return nil
}

func (s *PublishMetadataService) RunPublishWearableMetadata(csvFile string) error {
	s.l.Infof("preparing to parse metronion wearable metadata from CSV file %v", csvFile)
	metronionWearables, err := s.ParseMetronionWearableMetadata(csvFile)
	if err != nil {
		s.l.Errorw("error parse metronion wearable metadata")
		return err
	}
	s.l.Infow("parse metronion wearables metadata successfully", "metronion_wearables_length", len(metronionWearables))

	convertedData, err := convertMetronionWearableMetadata(metronionWearables)
	if err != nil {
		s.l.Errorw("error convert metronion wearables", "error", err)
		return err
	}
	s.l.Infow("convert metronion wearables metadata successfully", "metronion_wearables_length", len(convertedData))
	if err := s.persistentDb.SaveBatchWearableMetadata(convertedData); err != nil {
		s.l.Errorw("error save batch metronion wearables to persistent database", "error", err)
		return err
	}
	s.l.Infof("save batch metronion wearables successfully")

	return nil
}

func (s *PublishMetadataService) ParseMetronionMetadata(csvFile string) ([]MetronionMetadataCsv, error) {
	var result = make([]MetronionMetadataCsv, 0)

	metadataFile, err := os.OpenFile(csvFile, os.O_RDWR, os.ModePerm)
	if err != nil {
		s.l.Errorw("error open metadata file", "error", err)
		return nil, err
	}
	defer metadataFile.Close()

	if err != nil {
		s.l.Errorw("error read metronion metadata csv file", "error", err)
		return nil, err
	}

	if err := gocsv.UnmarshalFile(metadataFile, &result); err != nil {
		s.l.Errorw("error unmarshal csv file", "error", err)
		return nil, err
	}

	return result, nil
}

func (s *PublishMetadataService) ParseMetronionWearableMetadata(csvFile string) ([]MetronionWearableMetadata, error) {
	var result = make([]MetronionWearableMetadata, 0)

	metadataFile, err := os.OpenFile(csvFile, os.O_RDWR, os.ModePerm)
	if err != nil {
		s.l.Errorw("error open wearable metadata file", "error", err, "file", csvFile)
		return nil, err
	}
	defer metadataFile.Close()

	if err != nil {
		s.l.Errorw("error read metronion wearable metadata csv file", "error", err)
		return nil, err
	}

	if err := gocsv.UnmarshalFile(metadataFile, &result); err != nil {
		s.l.Errorw("error unmarshal csv file", "error", err)
		return nil, err
	}

	return result, nil
}

func convertMetronion(data []MetronionMetadataCsv) ([]metronion.MetronionMetadata, error) {
	if len(data) == 0 {
		return []metronion.MetronionMetadata{}, nil
	}
	var result = make([]metronion.MetronionMetadata, len(data))
	for i := range data {
		stats, err := convertStats(data[i].TotalStats)
		if err != nil {
			return nil, err
		}
		specialStats, err := convertStats(data[i].SpecialStats)
		if err != nil {
			return nil, err
		}
		wearableIDs := strings.Split(data[i].Wearables, ",")
		wearables := make([]metronion.Wearables, len(wearableIDs))
		for i := range wearableIDs {
			wearables[i] = metronion.Wearables{
				WearableId: wearableIDs[i],
			}
		}
		result[i] = metronion.MetronionMetadata{
			Id:           data[i].ID,
			Type:         strings.ToLower(string(data[i].Type)),
			Gender:       strings.ToLower(string(data[i].Gender)),
			Star:         uint64(data[i].Star),
			SpecialStar:  uint64(data[i].SpecialStar),
			BaseStats:    stats,
			SpecialStats: specialStats,
			Wearables:    wearables,
		}
	}
	return result, nil
}

func convertStats(data string) (services.Stats, error) {
	// convert stats
	statsTmp := strings.Split(data, ",")
	if len(statsTmp) != 5 {
		return services.Stats{}, fmt.Errorf("invalid stats data")
	}
	creativity, err := strconv.ParseUint(statsTmp[0], 10, 64)
	if err != nil {
		return services.Stats{}, fmt.Errorf("cannot parse `creativity` field")
	}
	charisma, err := strconv.ParseUint(statsTmp[1], 10, 64)
	if err != nil {
		return services.Stats{}, fmt.Errorf("cannot parse `charisma` field")
	}
	resolve, err := strconv.ParseUint(statsTmp[2], 10, 64)
	if err != nil {
		return services.Stats{}, fmt.Errorf("cannot parse `resolve` field")
	}
	fitness, err := strconv.ParseUint(statsTmp[3], 10, 64)
	if err != nil {
		return services.Stats{}, fmt.Errorf("cannot parse `resolve` field")
	}
	intellect, err := strconv.ParseUint(statsTmp[4], 10, 64)
	if err != nil {
		return services.Stats{}, fmt.Errorf("cannot parse `intellect` field")
	}

	stats := services.Stats{
		Creativity: creativity,
		Charisma:   charisma,
		Resolve:    resolve,
		Fitness:    fitness,
		Intellect:  intellect,
	}
	return stats, nil
}

func convertMetronionWearableMetadata(data []MetronionWearableMetadata) ([]metronion.Wearables, error) {
	if len(data) == 0 {
		return []metronion.Wearables{}, nil
	}
	var result = make([]metronion.Wearables, len(data))
	for i := range data {
		result[i] = metronion.Wearables{
			WearableId: data[i].ID,
			Type:       strings.ToLower(data[i].Type),
			Name:       data[i].Name,
			Gender:     strings.ToLower(string(data[i].Gender)),
			Rarity:     strings.ToLower(string(data[i].Rarity)),
			Stats: services.Stats{
				Creativity: data[i].Creativity,
				Charisma:   data[i].Charisma,
				Resolve:    data[i].Resolve,
				Fitness:    data[i].Fitness,
				Intellect:  data[i].Intellect,
			},
			IsRequired: data[i].Required,
			Changeable: data[i].Changeable,
			IsOrigin:   data[i].IsOrigin,
			Image:      composeWearableS3Uri(data[i].ID, data[i].Type),
		}
	}
	return result, nil
}

func composeWearableS3Uri(id string, wearableType string) string {
	tmp := strings.Split(id, "_")
	if len(tmp) < 2 {
		return fmt.Sprintf("%s/unknown.png", WEARABLE_S3_URI)
	}

	return fmt.Sprintf("%s/%s/%s/%v.png", WEARABLE_S3_URI, strings.ToLower(tmp[0]), strings.ToLower(wearableType), id)
}
