package main

import (
	"fmt"
	"log"

	"github.com/ariefsn/superhero-db/helper"
	"github.com/ariefsn/superhero-db/models"
	"github.com/ariefsn/superhero-db/services"
	"github.com/gocolly/colly"
)

const baseUrl = "https://www.superherodb.com"

func main() {
	sh := models.NewSuperheroModel()

	help := helper.Helper{}

	c := help.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// Profiles
	c.OnHTML("div.profile-titles", func(h *colly.HTMLElement) {
		h1 := h.ChildText("h1")
		h2 := h.ChildText("h2")
		h3 := h.ChildText("h3")
		portrait := baseUrl + h.ChildAttr("div.portrait > img", "src")

		sh.Name = h1
		sh.RealName = h2
		sh.Origin.Universe.Name = h3
		sh.Portrait = portrait
	})

	// Tabs
	c.OnHTML("ul.tab", func(h *colly.HTMLElement) {
		mainTabs := []string{"About", "History", "Powers", "Equipment", "Movies", "Gallery", "Battles"}

		checkMainTabs := func(title string) bool {
			for _, v := range mainTabs {
				if title == v {
					return true
				}
			}

			return false
		}

		h.ForEach("li", func(idx int, li *colly.HTMLElement) {
			title := li.ChildText("a > span")

			if checkMainTabs(title) {
				href := baseUrl + li.ChildAttr("a", "href")

				srv := new(services.Service)

				srv.BaseUrl = baseUrl
				srv.Href = href
				srv.Data = sh

				// fmt.Println(title, srv)

				switch title {
				case "About":
					services.About(srv)
				case "History":
					services.History(srv)
				case "Powers":
					services.Powers(srv)
				case "Equipment":
					services.EquipmentAndWeapons(srv)
				case "Gallery":
					services.Gallery(srv)
				}
			}
		})
	})

	c.OnScraped(func(res *colly.Response) {
		fmt.Println("Finished scrape:", res.Request.URL)

		// fmt.Println(toolkit.JsonStringIndent(sh, "\n"))
		help.WriteJsonFile(sh, "new")
	})

	// c.Visit(baseUrl + "/nick-fury/10-16352/")
	// c.Visit(baseUrl + "/nick-fury/10-326/")
	c.Visit(baseUrl + "/captain-america/10-12495/")

	c.Wait()
}
