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

func NewPowerStatsModel() *PowerStatsModel {
	m := new(PowerStatsModel)

	m.Intelligence = 0
	m.Strength = 0
	m.Speed = 0
	m.Durability = 0
	m.Power = 0
	m.Combat = 0
	m.Tier = 0

	return m
}
