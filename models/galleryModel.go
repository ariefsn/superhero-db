package models

import "strings"

type GalleryModel struct {
	Sm string `json:"sm" bson:"sm"`
	Md string `json:"md" bson:"md"`
	Lg string `json:"lg" bson:"lg"`
}

func NewGalleryModel(baseUrl, path string) *GalleryModel {
	sizes := map[string]string{
		"sm": "025",
		"md": "050",
		"lg": "075",
	}

	split := strings.Split(path, "/")

	size := split[2]

	return &GalleryModel{
		Sm: baseUrl + strings.Replace(path, size, sizes["sm"], 1),
		Md: baseUrl + strings.Replace(path, size, sizes["md"], 1),
		Lg: baseUrl + strings.Replace(path, size, sizes["lg"], 1),
	}
}
