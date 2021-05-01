package models

type SuperheroModel struct {
	Name        string          `json:"name" bson:"name"`
	RealName    string          `json:"realName" bson:"realName"`
	Portrait    string          `json:"portrait" bson:"portrait"`
	History     string          `json:"history" bson:"history"`
	PowerStats  PowerStatsModel `json:"powerStats" bson:"powerStats"`
	Origin      OriginModel     `json:"origin" bson:"origin"`
	Appearance  AppearanceModel `json:"appearance" bson:"appearance"`
	Connections ConnectionModel `json:"connections" bson:"connections"`
	Powers      PowerModel      `json:"powers" bson:"powers"`
	Item        ItemModel       `json:"item" bson:"item"`
	Gallery     []GalleryModel  `json:"gallery" bson:"gallery"`
}
