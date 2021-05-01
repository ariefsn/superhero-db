package models

type ConnectionModel struct {
	Occupation []string   `json:"occupation" bson:"occupation"`
	Base       string     `json:"base" bson:"base"`
	Teams      []UrlModel `json:"teams" bson:"teams"`
	Relatives  []string   `json:"relatives" bson:"relatives"`
}
