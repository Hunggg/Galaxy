package cockroachdb

import (
	"strings"

	"github.com/metrogalaxy-io/metrogalaxy-api/services"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/metronion"
)

type MetronionMetadata struct {
	Id           uint64 `gorm:"primaryKey;autoIncrement:false"`
	Gender       string
	Type         string
	Image        string
	Star         uint64
	SpecialStar  uint64
	BaseStats    services.Stats `gorm:"embedded;embeddedPrefix:stats_"`
	SpecialStats services.Stats `gorm:"embedded;embeddedPrefix:special_stats_"`
	Wearables    string         // list of metronion wearables id
}

func ConvertMetronionMetadata(inputData metronion.MetronionMetadata) MetronionMetadata {
	wearableIDs := make([]string, len(inputData.Wearables))
	for i := range inputData.Wearables {
		wearableIDs[i] = inputData.Wearables[i].WearableId
	}

	return MetronionMetadata{
		Id:           inputData.Id,
		Gender:       inputData.Gender,
		Type:         inputData.Type,
		Image:        inputData.Image,
		Star:         inputData.Star,
		SpecialStar:  inputData.SpecialStar,
		BaseStats:    inputData.BaseStats,
		SpecialStats: inputData.SpecialStats,
		Wearables:    strings.Join(wearableIDs, ","),
	}
}

type Wearables struct {
	ID         uint64 `gorm:"primaryKey"`
	WearableId string `gorm:"unique"`
	Type       string
	Name       string
	Gender     string
	Rarity     string
	Changeable bool
	IsOrigin   bool
	IsRequired bool
	Image      string
	Stats      services.Stats `gorm:"embedded"`
}

type WearablesModelDB struct {
	ID         uint64 `gorm:"primaryKey"`
	WearableId string `gorm:"unique"`
	Type       string
	Name       string
	Gender     string
	Rarity     string
	Changeable bool
	IsOrigin   bool
	IsRequired bool
	Image      string
	Creativity uint64
	Charisma   uint64
	Resolve    uint64
	Fitness    uint64
	Intellect  uint64
}

func ConvertWearables(inputData metronion.Wearables) Wearables {
	return Wearables{
		WearableId: inputData.WearableId,
		Type:       inputData.Type,
		Name:       inputData.Name,
		Gender:     inputData.Gender,
		Rarity:     inputData.Rarity,
		Changeable: inputData.Changeable,
		IsOrigin:   inputData.IsOrigin,
		IsRequired: inputData.IsRequired,
		Image:      inputData.Image,
		Stats:      inputData.Stats,
	}
}

type Metronion struct {
	Id                     uint64
	Name                   string
	VersionId              uint64
	CreatedAtTimestamp     uint64
	UpdatedAtTimestamp     uint64
	CreatedAtBlock         uint64
	UpdatedAtBlock         uint64
	Owner                  string
	LastPrice              float64
	CurrentPrice           float64
	TopBid                 float64
	ListedBy               string
	Image                  string
	Gender                 string
	Type                   string
	Star                   uint64
	SpecialStar            uint64
	StatsCreativity        uint64
	StatsCharisma          uint64
	StatsResolve           uint64
	StatsFitness           uint64
	StatsIntellect         uint64
	SpecialStatsCreativity uint64
	SpecialStatsCharisma   uint64
	SpecialStatsResolve    uint64
	SpecialStatsFitness    uint64
	SpecialStatsIntellect  uint64
	Wearables              string // 0,1,2,3,4
}

// type Metronion struct {
// 	Id                 uint64 `gorm:"primaryKey;autoIncrement:false"`
// 	Name               string
// 	VersionId          uint64
// 	CreatedAtTimestamp uint64
// 	UpdatedAtTimestamp uint64
// 	CreatedAtBlock     uint64
// 	UpdatedAtBlock     uint64
// 	Owner              string
// 	Wearables          string // list of metronion wearables id
// }
//
//func ConvertMetronion(inputData metronion.Metronion) Metronion {
//	wearableIDs := make([]string, len(inputData.Wearables))
//	for i := range inputData.Wearables {
//		wearableIDs[i] = inputData.Wearables[i].WearableId
//	}
//
//	return Metronion{
//		Id:                 inputData.Id,
//		Name:               inputData.Name,
//		VersionId:          inputData.VersionId,
//		CreatedAtTimestamp: inputData.CreatedAtTimestamp,
//		UpdatedAtTimestamp: inputData.UpdatedAtTimestamp,
//		CreatedAtBlock:     inputData.CreatedAtBlock,
//		UpdatedAtBlock:     inputData.UpdatedAtBlock,
//		Owner:              inputData.Owner,
//		Wearables:          strings.Join(wearableIDs, ","),
//	}
//}