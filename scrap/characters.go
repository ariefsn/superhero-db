package scrap

import (
	"fmt"
	"log"
	"strings"

	"github.com/ariefsn/superhero-db/helper"
	"github.com/ariefsn/superhero-db/models"
	"github.com/gocolly/colly"
)

func getCharList(url string) []*models.CharacterModel {
	chars := []*models.CharacterModel{}

	help := helper.Helper{}

	c := help.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("ul.list.list-md", func(h *colly.HTMLElement) {

		h.ForEach("li", func(iLi int, li *colly.HTMLElement) {
			char := models.NewCharacterModel()

			char.RealName = li.ChildText("span.suffix.level-1")
			char.Universe = li.ChildText("span.suffix.level-2")
			char.Path = li.ChildAttr("a", "href")
			char.Name = strings.Replace(li.Text, char.RealName, "", 1)
			char.Name = strings.Replace(char.Name, char.Universe, "", 1)

			chars = append(chars, char)
		})
	})

	c.Visit(url)

	c.Wait()

	return chars
}

func GetCharacters(baseUrl string) []*models.CharacterModel {
	baseUrl += "/characters/"

	chars := []*models.CharacterModel{}

	totalPage := 70

	for page := 0; page < totalPage; page++ {
		url := fmt.Sprintf("%s?page_nr=%v", baseUrl, (page + 1))

		chars = append(chars, getCharList(url)...)
	}

	return chars
}
