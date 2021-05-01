package services

type Service struct {
	BaseUrl string
	Href    string
}

func (s *Service) FullPath() string {
	return s.BaseUrl + s.Href
}

func (s *Service) separator(sep string) string {
	return "***" + sep + "***"
}
