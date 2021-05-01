package services

import (
	"strings"

	"github.com/gocolly/colly"
)

func History(s *Service) {
	c := colly.NewCollector()

	c.OnHTML(".column.col-12.text-columns-2", func(h *colly.HTMLElement) {
		t := h.ChildText("h3")
		_ = strings.TrimSpace(strings.ReplaceAll(h.Text, t, ""))
		// fmt.Println(hist)
	})

	c.Visit(s.Href)
}
