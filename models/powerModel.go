package models

type PowerDetailsModel struct {
	Title       string
	Description string
}

type PowerModel struct {
	Summary string
	Details []PowerDetailsModel
}
