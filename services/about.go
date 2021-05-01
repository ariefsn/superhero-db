package services

import (
	"strings"

	"github.com/ariefsn/superhero-db/models"
	"github.com/gocolly/colly"
	"github.com/novalagung/gubrak"
)

func About(s *Service) {
	c := colly.NewCollector()

	selectedTable := ""

	clearList := func(list []string) []string {
		r := gubrak.From(list).Map(func(each string, i int) string {
			return strings.TrimSpace(each)
		}).Result()

		r = gubrak.From(r).Filter(func(each string, i int) bool {
			return each != ""
		}).Result()

		return r.([]string)
	}

	c.OnHTML("table.profile-table", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(iTr int, tr *colly.HTMLElement) {

			tr.ForEach("td", func(iTd int, td *colly.HTMLElement) {
				if iTr == 0 && iTd == 0 {
					if td.Text == "Creator" {
						selectedTable = "Origin"
					} else if td.Text == "Occupation" {
						selectedTable = "Connections"
					} else {
						selectedTable = "Appearance"
					}
				}

				// get value from 2nd td
				if iTd == 1 {
					if selectedTable == "Origin" {
						switch iTr {
						case 0:
							s.Data.Origin.Creator.Name = td.Text
							s.Data.Origin.Creator.Url = s.BaseUrl + td.ChildAttr("a", "href")
						case 1:
							s.Data.Origin.Universe.Name = td.Text
							s.Data.Origin.Universe.Url = s.BaseUrl + td.ChildAttr("a", "href")
						case 2:
							s.Data.Origin.FullName = td.Text
						case 3:
							s.Data.Origin.AlterEgos = strings.TrimSpace(td.Text)
						case 4:
							s.Data.Origin.Aliases = strings.Split(td.Text, ", ")
						case 5:
							s.Data.Origin.PlaceOfBirth = td.Text
						case 6:
							s.Data.Origin.FirstAppearance = td.Text
						case 7:
							s.Data.Origin.Alignment = td.Text
						}
					} else if selectedTable == "Connections" {
						switch iTr {
						case 0:
							s.Data.Connections.Occupation = clearList(strings.Split(td.Text, ";"))
						case 1:
							s.Data.Connections.Base = strings.TrimSpace(td.Text)
						case 2:
							teams := []models.UrlModel{}

							td.ForEach("a", func(iA int, a *colly.HTMLElement) {
								teams = append(teams, models.UrlModel{
									Name: a.Text,
									Url:  s.BaseUrl + a.Attr("href"),
								})
							})

							s.Data.Connections.Teams = teams
						case 3:
							s.Data.Connections.Relatives = clearList(strings.Split(td.Text, ";"))
						}
					}
				}
			})
		})
	})

	c.Visit(s.Href)
}
