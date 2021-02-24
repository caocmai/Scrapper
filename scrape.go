package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

// FoodList stores a list of food
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
	c := colly.NewCollector(
		// visit only domains of calories.info or www.calories.info
		colly.AllowedDomains("calories.info", "www.calories.info"),
		// caching to prevent multiple download of pages
		colly.CacheDir("./calorie_cache"),
	)

	cloned := c.Clone()

	cloned.OnHTML("body", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)
		fmt.Println("went inside of here")
	})

	// On every a element which has specified attribute call callback
	c.OnHTML(`body`, func(e *colly.HTMLElement) {
		// fmt.Println(e.Text)
		tmpFoodList := FoodList{}
		link := e.Attr("href")
		fmt.Println(link)

		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			// fmt.Println(e.Text)

			tmpFood := Food{}

			tmpFood.Name = el.ChildText("td:first-child")
			tmpFood.Calorie = el.ChildText("td:nth-child(5)")

			tmpFoodList.Foods = append(tmpFoodList.Foods, tmpFood)

		})

		js, err := json.MarshalIndent(tmpFoodList, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(js))

		_ = ioutil.WriteFile("foods.json", js, 0644)

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping
	c.Wait()
	c.Visit("https://www.calories.info/food/meat")
}
