package publishmetadata

// type MetronionPartType string

// const (
// 	BodyPart            MetronionPartType = "Body"
// 	HeadPart            MetronionPartType = "Head"
// 	TopPart             MetronionPartType = "Top"
// 	BottomPart          MetronionPartType = "Bottom"
// 	ShoesPart           MetronionPartType = "Shoes"
// 	FacePart            MetronionPartType = "Face"
// 	FaceAccessoriesPart MetronionPartType = "Face_Accessories"
// 	HairPart            MetronionPartType = "Hair"
// 	HairLongPart        MetronionPartType = "Hair_Long"
// 	SetPart             MetronionPartType = "Set"
// )

type MetronionMetadataCsv struct {
	ID           uint64 `csv:"ID"`
	Type         string `csv:"Type"`
	Gender       string `csv:"Gender"`
	TotalStats   string `csv:"TotalStats"`
	SpecialStats string `csv:"SpecialStats"`
	Star         int    `csv:"Star"`
	SpecialStar  int    `csv:"SpecialStar"`
	Wearables    string `csv:"Parts"`
}

type MetronionWearableMetadata struct {
	ID         string `csv:"ID"`
	Type       string `csv:"Type"`
	Name       string `csv:"Name"`
	Gender     string `csv:"Gender"`
	Rarity     string `csv:"Rarity"`
	Creativity uint64 `csv:"Creativity"`
	Charisma   uint64 `csv:"Charisma"`
	Resolve    uint64 `csv:"Resolve"`
	Fitness    uint64 `csv:"Fitness"`
	Intellect  uint64 `csv:"Intellect"`
	Required   bool   `csv:"Required"`
	Changeable bool   `csv:"Changeable"`
	IsOrigin   bool   `csv:"IsOrigin"`
}
