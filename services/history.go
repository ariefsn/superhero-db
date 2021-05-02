package services

import (
	"fmt"
	"strings"

	"github.com/ariefsn/superhero-db/helper"
	"github.com/gocolly/colly"
)

func History(s *Service) {
	help := helper.Helper{}

	c := help.NewCollector()

	c.OnHTML(".column.col-12.text-columns-2", func(h *colly.HTMLElement) {
		t := h.ChildText("h3")
		hist := strings.TrimSpace(strings.ReplaceAll(h.Text, t, ""))
		s.Data.History = hist
	})

	c.OnScraped(func(res *colly.Response) {
		fmt.Println("Finished scrape:", res.Request.URL)
	})

	c.Visit(s.Href)

	c.Wait()
}
