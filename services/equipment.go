package services

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/ariefsn/superhero-db/models"
	"github.com/eaciit/toolkit"
	"github.com/gocolly/colly"
)

func EquipmentAndWeapons(s *Service) {
	c := colly.NewCollector()

	// eq := new(Equipment)

	c.OnHTML(".column.col-8.col-md-12", func(h *colly.HTMLElement) {
		key := ""

		listH3 := []string{}
		listH4 := []string{}
		mEq := map[string][]string{}
		h4Prefix := "#h4#"

		h4WithPrefix := func(h4 string) string {
			return h4Prefix + h4 + h4Prefix
		}

		getKey := func(title string) string {
			if strings.Contains(title, "weapon") {
				return "Weapons"
			}
			return "Equipment"
		}

		h.DOM.Children().Each(func(i int, g *goquery.Selection) {
			tagName := goquery.NodeName(g)

			if tagName == "h3" {
				key = getKey(g.Text())

				listH3 = append(listH3, g.Text())
			} else if tagName == "h4" {
				_list := append(mEq[key], g.Text())
				mEq[key] = _list

				g.SetText(h4WithPrefix((g.Text())))

				listH4 = append(listH4, g.Text())
			}
		})

		allText := strings.TrimSpace(h.DOM.Text())

		splitText := map[string]string{}

		res := new(models.EquipmentModel)

		for iH3, h3 := range listH3 {
			h3Separator := s.separator("h3")

			h3text := strings.TrimSpace(strings.Replace(allText, h3, h3Separator, 1))

			if iH3 < len(listH3)-1 {
				h3text = strings.TrimSpace(strings.Replace(h3text, listH3[iH3+1], h3Separator, 1))
			}

			splitByH3 := strings.Split(h3text, h3Separator)
			if len(splitByH3) > 1 {
				content := strings.TrimSpace(splitByH3[1])

				_k := getKey(h3)

				splitText[_k] = content
			}
		}

		for k, _text := range splitText {
			_list := mEq[k]

			for iH4, h4 := range _list {
				h4Separator := s.separator("h4")

				h4text := strings.TrimSpace(strings.Replace(_text, h4WithPrefix(h4), h4Separator, 1))

				if iH4 < len(_list)-1 {
					h4text = strings.TrimSpace(strings.Replace(h4text, h4WithPrefix(_list[iH4+1]), h4Separator, 1))
				}

				splitByH4 := strings.Split(h4text, h4Separator)
				if len(splitByH4) > 1 {
					content := strings.TrimSpace(splitByH4[1])

					newData := models.EquipmentItemModel{
						Title:       h4,
						Description: content,
					}

					if k == "Weapons" {
						res.Weapon.Items = append(res.Weapon.Items, newData)
					} else {
						res.Equipment.Items = append(res.Equipment.Items, newData)
					}
				}
			}
		}

		fmt.Println(toolkit.JsonStringIndent(res, "n"))

		// for i, v := range split {
		// 	if i == 0 {
		// 		res.Summary = split[i]
		// 	} else {
		// 		res.Details = append(res.Details, PowerDetails{
		// 			Title:       powerTitles[i-1],
		// 			Description: v,
		// 		})
		// 	}
		// }

		// fmt.Println("res =>", res)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("equip req:", r.URL)
	})

	c.Visit(s.Href)
}
