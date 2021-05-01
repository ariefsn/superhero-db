package models

type PowerStatsModel struct {
	Intelligence int `json:"intelligence" bson:"intelligence"`
	Strength     int `json:"strength" bson:"strength"`
	Speed        int `json:"speed" bson:"speed"`
	Durability   int `json:"durability" bson:"durability"`
	Power        int `json:"power" bson:"power"`
	Combat       int `json:"combat" bson:"combat"`
	Tier         int `json:"tier" bson:"tier"`
}
