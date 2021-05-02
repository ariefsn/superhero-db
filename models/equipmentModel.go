package models

type EquipmentItemModel struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}

type EquipmentDetailsModel struct {
	Summary string               `json:"summary" bson:"summary"`
	List    []EquipmentItemModel `json:"list" bson:"list"`
}

type ItemModel struct {
	Equipment EquipmentDetailsModel `json:"equipment" bson:"equipment"`
	Weapon    EquipmentDetailsModel `json:"weapon" bson:"weapon"`
}

func NewEquipmentDetailsModel() *EquipmentDetailsModel {
	m := new(EquipmentDetailsModel)

	m.List = []EquipmentItemModel{}

	return m
}

func NewItemModel() *ItemModel {
	m := new(ItemModel)

	m.Equipment = *NewEquipmentDetailsModel()
	m.Weapon = *NewEquipmentDetailsModel()

	return m
}
