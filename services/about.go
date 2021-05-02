package services

import (
	"fmt"
	"strings"

	"github.com/ariefsn/superhero-db/helper"
	"github.com/ariefsn/superhero-db/models"
	"github.com/eaciit/toolkit"
	"github.com/gocolly/colly"
)

func About(s *Service) {
	help := helper.Helper{}

	c := help.NewCollector()

	selectedTable := ""

	c.OnHTML(".column.col-12", func(h *colly.HTMLElement) {
		h3 := h.ChildText("h3")

		if strings.TrimSpace(strings.ToLower(h3)) == "super powers" {
			h.ForEach("a", func(iA int, a *colly.HTMLElement) {
				s.Data.SuperPower = append(s.Data.SuperPower, models.UrlModel{
					Name: a.Text,
					Url:  s.BaseUrl + a.Attr("href"),
				})
			})
		}
	})

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
							alterEgos := []models.AlterEgosModel{}

							td.ForEach(".shdbcard3.cat-10.card-xs", func(iAe int, ae *colly.HTMLElement) {
								alterEgos = append(alterEgos, models.AlterEgosModel{
									Url:      ae.ChildAttr("a", "href"),
									Class:    ae.ChildText(".shdbclass span"),
									Verse:    ae.ChildText(".verse span"),
									Name:     ae.ChildText(".name"),
									RealName: ae.ChildText(".realname"),
									Image:    s.BaseUrl + ae.ChildAttr("div.image img", "src"),
								})
							})

							s.Data.Origin.AlterEgos = alterEgos
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
							s.Data.Connections.Occupation = help.ClearList(strings.Split(td.Text, ";"))
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
							s.Data.Connections.Relatives = help.ClearList(strings.Split(td.Text, ";"))
						}
					} else if selectedTable == "Appearance" {
						switch iTr {
						case 0:
							s.Data.Appearance.Gender = strings.TrimSpace(td.Text)
						case 1:
							s.Data.Appearance.Type.Name = td.Text
							s.Data.Appearance.Type.Url = s.BaseUrl + td.ChildAttr("a", "href")
						case 2:
							s.Data.Appearance.Height = strings.Replace(help.ExtractText(strings.TrimSpace(td.Text)), "  ", " / ", 1)
						case 3:
							s.Data.Appearance.Weight = strings.Replace(help.ExtractText(strings.TrimSpace(td.Text)), "  ", " / ", 1)
						case 4:
							s.Data.Appearance.EyeColor = strings.TrimSpace(td.Text)
						case 5:
							s.Data.Appearance.HairColor = strings.TrimSpace(td.Text)
						}
					}
				}
			})
		})
	})

	c.OnHTML(".stat-bar", func(h *colly.HTMLElement) {
		val := toolkit.ToInt(h.ChildText(".stat-value"), "")

		switch strings.ToLower(h.ChildText("label")) {
		case "intelligence":
			s.Data.PowerStats.Intelligence = val
		case "stregth":
			s.Data.PowerStats.Strength = val
		case "speed":
			s.Data.PowerStats.Speed = val
		case "durability":
			s.Data.PowerStats.Durability = val
		case "power":
			s.Data.PowerStats.Power = val
		case "combat":
			s.Data.PowerStats.Intelligence = val
		case "tier":
			s.Data.PowerStats.Intelligence = val
		}
	})

	c.OnScraped(func(res *colly.Response) {
		fmt.Println("Finished scrape:", res.Request.URL)
	})

	c.Visit(s.Href)

	c.Wait()
}
