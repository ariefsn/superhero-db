package models

type OriginModel struct {
	Creator         UrlModel
	Universe        UrlModel
	FullName        string
	AlterEgos       string
	Aliases         []string
	PlaceOfBirth    string
	FirstAppearance string
	Alignment       string
}
