package models

type AppearanceModel struct {
	Gender    string   `json:"gender" bson:"gender"`
	Type      UrlModel `json:"type" bson:"type"`
	Height    string   `json:"height" bson:"height"`
	Weight    string   `json:"weight" bson:"weight"`
	EyeColor  string   `json:"eyeColor" bson:"eyeColor"`
	HairColor string   `json:"hairColor" bson:"hairColor"`
}
