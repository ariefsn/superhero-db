package models

type OriginModel struct {
	Creator         UrlModel `json:"creator" bson:"creator"`
	Universe        UrlModel `json:"universe" bson:"universe"`
	FullName        string   `json:"fullName" bson:"fullName"`
	AlterEgos       string   `json:"alterEgos" bson:"alterEgos"`
	Aliases         []string `json:"aliases" bson:"aliases"`
	PlaceOfBirth    string   `json:"placeOfBirth" bson:"placeOfBirth"`
	FirstAppearance string   `json:"firstAppearance" bson:"firstAppearance"`
	Alignment       string   `json:"alignment" bson:"alignment"`
}
