package services

import (
	"github.com/ariefsn/superhero-db/models"
	"github.com/gocolly/colly"
)

func Gallery(s *Service) {
	c := colly.NewCollector()

	c.OnHTML("img.gallery-thumb", func(h *colly.HTMLElement) {
		src := h.Attr("src")

		s.Data.Gallery = append(s.Data.Gallery, *models.NewGalleryModel(s.BaseUrl, src))
	})

	c.Visit(s.Href)
}
