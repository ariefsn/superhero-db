package main

import (
	"fmt"
	"log"

	"github.com/ariefsn/superhero-db/services"
	"github.com/gocolly/colly"
)

const baseUrl = "https://www.superherodb.com"

func main() {

	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// Profiles
	c.OnHTML("div.profile-titles", func(h *colly.HTMLElement) {
		h1 := h.ChildText("h1")
		h2 := h.ChildText("h2")
		h3 := h.ChildText("h3")
		portrait := baseUrl + h.ChildAttr("div.portrait > img", "src")

		fmt.Println("h1", h1)
		fmt.Println("h2", h2)
		fmt.Println("h3", h3)
		fmt.Println("portrait", portrait)
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

				fmt.Println(idx, title, href)

				srv := new(services.Service)

				srv.BaseUrl = baseUrl
				srv.Href = href

				// fmt.Println(title, srv)

				switch title {
				case "History":
					services.History(srv)
				case "Powers":
					services.Powers(srv)
				case "Equipment":
					services.EquipmentAndWeapons(srv)
				}
			}
		})
	})

	// // About
	// c.OnHTML("div.profile-titles", func(h *colly.HTMLElement) {
	// })

	c.OnScraped(func(res *colly.Response) {
		fmt.Println("Finished scrape:", res.Request.URL)
	})

	c.Visit(baseUrl + "/nick-fury/10-16352/")
}
