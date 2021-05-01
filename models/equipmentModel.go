package models

type EquipmentItemModel struct {
	Title       string
	Description string
}

type EquipmentDetailsModel struct {
	Summary string
	Items   []EquipmentItemModel
}

type ItemModel struct {
	Equipment EquipmentDetailsModel
	Weapon    EquipmentDetailsModel
}
