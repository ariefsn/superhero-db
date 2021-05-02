package models

type CharacterModel struct {
	Name     string `json:"name" bson:"name"`
	RealName string `json:"realName" bson:"realName"`
	Universe string `json:"universe" bson:"universe"`
	Path     string `json:"path" bson:"path"`
}

func NewCharacterModel() *CharacterModel {
	m := new(CharacterModel)

	return m
}
