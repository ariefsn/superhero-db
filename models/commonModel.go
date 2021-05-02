package models

type UrlModel struct {
	Name string `json:"name" bson:"name"`
	Url  string `json:"url" bson:"url"`
}

func NewUrlModel() *UrlModel {
	m := new(UrlModel)

	return m
}
