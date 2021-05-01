package services

import "github.com/ariefsn/superhero-db/models"

type Service struct {
	BaseUrl string
	Href    string
	Data    *models.SuperheroModel
}

func (s *Service) FullPath() string {
	return s.BaseUrl + s.Href
}

func (s *Service) separator(sep string) string {
	return "***" + sep + "***"
}
