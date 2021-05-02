package services

import (
	"fmt"

	"github.com/ariefsn/superhero-db/helper"
	"github.com/ariefsn/superhero-db/models"
	"github.com/gocolly/colly"
)

func Gallery(s *Service) {
	help := helper.Helper{}

	c := help.NewCollector()

	c.OnHTML("img.gallery-thumb", func(h *colly.HTMLElement) {
		src := h.Attr("src")

		s.Data.Gallery = append(s.Data.Gallery, *models.NewGalleryModel(s.BaseUrl, src))
	})

	c.OnScraped(func(res *colly.Response) {
		fmt.Println("Finished scrape:", res.Request.URL)
	})

	c.Visit(s.Href)

	c.Wait()
}
