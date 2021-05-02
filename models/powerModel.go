package models

type PowerDetailsModel struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}

type PowerModel struct {
	Summary string              `json:"summary" bson:"summary"`
	Details []PowerDetailsModel `json:"details" bson:"details"`
}

func NewPowerModel() *PowerModel {
	m := new(PowerModel)

	m.Details = []PowerDetailsModel{}

	return m
}
