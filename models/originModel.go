package models

type AlterEgosModel struct {
	Url      string `json:"url" bson:"url"`
	Name     string `json:"name" bson:"name"`
	Class    string `json:"class" bson:"class"`
	Verse    string `json:"verse" bson:"verse"`
	RealName string `json:"realName" bson:"realName"`
	Image    string `json:"image" bson:"image"`
}

type OriginModel struct {
	Creator         UrlModel         `json:"creator" bson:"creator"`
	Universe        UrlModel         `json:"universe" bson:"universe"`
	FullName        string           `json:"fullName" bson:"fullName"`
	AlterEgos       []AlterEgosModel `json:"alterEgos" bson:"alterEgos"`
	Aliases         []string         `json:"aliases" bson:"aliases"`
	PlaceOfBirth    string           `json:"placeOfBirth" bson:"placeOfBirth"`
	FirstAppearance string           `json:"firstAppearance" bson:"firstAppearance"`
	Alignment       string           `json:"alignment" bson:"alignment"`
}

func NewOriginModel() *OriginModel {
	m := new(OriginModel)

	m.Creator = *NewUrlModel()
	m.Universe = *NewUrlModel()
	m.Aliases = []string{}
	m.AlterEgos = []AlterEgosModel{}

	return m
}
