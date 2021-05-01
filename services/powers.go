package services

import (
	"fmt"
	"strings"

	"github.com/ariefsn/superhero-db/models"
	"github.com/gocolly/colly"
)

func Powers(s *Service) {
	c := colly.NewCollector()

	c.OnHTML(".column.col-8.col-md-12", func(h *colly.HTMLElement) {
		t := h.ChildText("h3")
		powers := strings.TrimSpace(strings.ReplaceAll(h.Text, t, ""))

		powerTitles := []string{}

		h.ForEach("h4", func(i int, h4 *colly.HTMLElement) {
			powerTitle := h4.Text

			powerTitles = append(powerTitles, powerTitle)

			powers = strings.TrimSpace(strings.ReplaceAll(powers, powerTitle, "***"))
		})

		split := strings.Split(strings.TrimSpace(powers), "***")

		res := new(models.PowerModel)

		for i, v := range split {
			if i == 0 {
				res.Summary = split[i]
			} else {
				res.Details = append(res.Details, models.PowerDetailsModel{
					Title:       powerTitles[i-1],
					Description: v,
				})
			}
		}

		s.Data.Powers = *res
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("powers req:", r.URL, s.Href)
	})

	c.Visit(s.Href)
}
