package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Food stores calorie info of a food
type Food struct {
	Name    string `json: "name"`
	Calorie int    `json: "calorie"`
}

func main() {
	// default collector
	c := colly.NewCollector()

	// On every a element which has specified attribute call callback
	c.OnHTML(".table", func(e *colly.HTMLElement) {
		// fmt.Println(e.Text)

		// link := e.Attr("href")
		// fmt.Println(e.ChildText("td.food.sorting_1"))
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			// fmt.Println(e.Text)

			fmt.Println(el.ChildText("td:first-child"))

			fmt.Println(el.ChildText("td:nth-child(5)"))

			// fmt.Println(e.ChildText(".food > a"))

		})

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping
	c.Visit("https://www.calories.info/food/meat")
}
