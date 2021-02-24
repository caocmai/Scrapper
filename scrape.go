package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type FoodList struct {
	Foods []Food
}

// Food stores calorie info of a food
type Food struct {
	Name    string `json: "name"`
	Calorie string `json: "calorie"`
}

func main() {
	// default collector
	c := colly.NewCollector()

	// On every a element which has specified attribute call callback
	c.OnHTML(".table", func(e *colly.HTMLElement) {
		// fmt.Println(e.Text)
		tmpFoodList := FoodList{}
		// link := e.Attr("href")
		// fmt.Println(e.ChildText("td.food.sorting_1"))
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			// fmt.Println(e.Text)

			tmpFood := Food{}

			tmpFood.Name = el.ChildText("td:first-child")

			tmpFood.Calorie = el.ChildText("td:nth-child(5)")

			// fmt.Println(e.ChildText(".food > a"))
			tmpFoodList.Foods = append(tmpFoodList.Foods, tmpFood)

		})

		js, err := json.MarshalIndent(tmpFoodList, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(js))

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping
	c.Visit("https://www.calories.info/food/meat")
}
