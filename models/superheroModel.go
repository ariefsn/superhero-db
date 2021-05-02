package models

type SuperheroModel struct {
	Name        string          `json:"name" bson:"name"`
	RealName    string          `json:"realName" bson:"realName"`
	Portrait    string          `json:"portrait" bson:"portrait"`
	History     string          `json:"history" bson:"history"`
	PowerStats  PowerStatsModel `json:"powerStats" bson:"powerStats"`
	SuperPower  []UrlModel      `json:"superPower" bson:"superPower"`
	Origin      OriginModel     `json:"origin" bson:"origin"`
	Appearance  AppearanceModel `json:"appearance" bson:"appearance"`
	Connections ConnectionModel `json:"connections" bson:"connections"`
	Powers      PowerModel      `json:"powers" bson:"powers"`
	Item        ItemModel       `json:"item" bson:"item"`
	Gallery     []GalleryModel  `json:"gallery" bson:"gallery"`
}

func NewSuperheroModel() *SuperheroModel {
	m := new(SuperheroModel)

	m.PowerStats = *NewPowerStatsModel()
	m.Origin = *NewOriginModel()
	m.Appearance = *NewAppearanceModel()
	m.Connections = *NewConnectionModel()
	m.Powers = *NewPowerModel()
	m.Item = *NewItemModel()
	m.Gallery = []GalleryModel{}
	m.SuperPower = []UrlModel{}

	return m
}
