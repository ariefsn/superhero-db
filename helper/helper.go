package helper

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/novalagung/gubrak"
)

type Helper struct{}

func (h *Helper) NewCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		RandomDelay: 1 * time.Second,
	})

	return c
}

func (h *Helper) ClearList(list []string) []string {
	r := gubrak.From(list).Map(func(each string, i int) string {
		return strings.TrimSpace(each)
	}).Result()

	r = gubrak.From(r).Filter(func(each string, i int) bool {
		return each != ""
	}).Result()

	return r.([]string)
}

func (h *Helper) ExtractText(text string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9 .'\"]+")
	return re.ReplaceAllString(text, "")
}

func (h *Helper) WriteJsonFile(data interface{}, fileName string) {
	file, _ := json.MarshalIndent(data, "", "\t")

	if fileName == "" {
		fileName = "test"
	}

	_ = ioutil.WriteFile(fileName+".json", file, 0644)
}
