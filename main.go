package main

import (
	"strings"

	"github.com/ariefsn/superhero-db/helper"
	"github.com/ariefsn/superhero-db/models"
	"github.com/ariefsn/superhero-db/scrap"
)

const baseUrl = "https://www.superherodb.com"

func main() {
	help := helper.Helper{}

	chars := scrap.GetCharacters(baseUrl)

	help.WriteJsonFile(map[string]interface{}{
		"data": chars,
	}, "", "characters")

	all := []*models.SuperheroModel{}

	for _, c := range chars {
		removed := strings.TrimPrefix(c.Path, "/")
		removed = strings.TrimSuffix(removed, "/")
		split := strings.Split(removed, "/")

		sh := scrap.GetCharacterDetails(baseUrl, c.Path, split[0], split[1])

		all = append(all, sh)
	}

	help.WriteJsonFile(map[string]interface{}{
		"data": all,
	}, "", "allCharacters")
}
