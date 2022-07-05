package services

import "time"

type Stats struct {
	Creativity uint64
	Charisma   uint64
	Resolve    uint64
	Fitness    uint64
	Intellect  uint64
}

type OnChainBaseActivity struct {
	From        string
	To          string
	Timestamp   uint64
	BlockNumber uint64
	TxHash      string
}

type MetronionWearableType string

const (
	MetronionWearableTypeHair               MetronionWearableType = "hair"
	MetronionWearableTypeSkin               MetronionWearableType = "skin"
	MetronionWearableTypeFace               MetronionWearableType = "face"
	MetronionWearableTypeFaceAccessories    MetronionWearableType = "face_accessories"
	MetronionWearableTypeTop                MetronionWearableType = "top"
	MetronionWearableTypeBottom             MetronionWearableType = "bottom"
	MetronionWearableTypeShoes              MetronionWearableType = "shoes"
	MetronionWearableTypeCoat               MetronionWearableType = "coat"
	MetronionWearableTypeLeftHandEquipment  MetronionWearableType = "left_hand_equipment"
	MetronionWearableTypeRightHandEquipment MetronionWearableType = "right_hand_equipment"
	MetronionWearableTypeVehicles           MetronionWearableType = "vehicles"
	MetronionWearableTypeWing               MetronionWearableType = "wing"
	MetronionWearableTypeTail               MetronionWearableType = "tail"
	MetronionWearableTypePet                MetronionWearableType = "pet"
	MetronionWearableTypeSet                MetronionWearableType = "set"
	MetronionWearableTypeHairLong           MetronionWearableType = "hair_long"
	MetronionWearableTypeHeadSkin           MetronionWearableType = "head_skin"
)

const (
	RarityCommon    string = "common"
	RarityUncommon  string = "uncommon"
	RarityRare      string = "rare"
	RarityEpic      string = "epic"
	RarityLegendary string = "legendary"
	RarityMythical  string = "mythical"
)

var MapWearablesType = map[uint64]string{
	0:  string(MetronionWearableTypeHair),
	1:  string(MetronionWearableTypeFaceAccessories),
	2:  string(MetronionWearableTypeTop),
	3:  string(MetronionWearableTypeBottom),
	4:  string(MetronionWearableTypeShoes),
	5:  string(MetronionWearableTypeCoat),
	6:  string(MetronionWearableTypeLeftHandEquipment),
	7:  string(MetronionWearableTypeRightHandEquipment),
	8:  string(MetronionWearableTypeVehicles),
	9:  string(MetronionWearableTypeWing),
	10: string(MetronionWearableTypeTail),
	11: string(MetronionWearableTypePet),
	12: string(MetronionWearableTypeSet),
}

var MapRarity = map[uint64]string{
	0: RarityCommon,
	1: RarityUncommon,
	2: RarityRare,
	3: RarityEpic,
	4: RarityLegendary,
	5: RarityMythical,
}

const (
	DefaultCockroachDbTimeout = 30 * time.Second
)
