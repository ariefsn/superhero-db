package models

type SuperheroModel struct {
	Name        string
	RealName    string
	Portrait    string
	History     string
	PowerStats  PowerStatsModel
	Origin      OriginModel
	Appearance  AppearanceModel
	Connections ConnectionModel
	Powers      PowerModel
	Item        ItemModel
	Gallery     []GalleryModel
}
