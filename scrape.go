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
	FoodType string
	Foods    []Food
}

// Food stores calorie info of a food
type Food struct {
	Name    string `json: "name"`
	Calorie string `json: "calorie"`
}

func main() {
	// Default collector
	c := colly.NewCollector(
		// Visit only domains of calories.info or www.calories.info
		colly.AllowedDomains("calories.info", "www.calories.info"),
		// Caching to prevent multiple download of pages
		colly.CacheDir("./calorie_cache"),
	)

	// So can do two things at once
	getFoodInfo := c.Clone()

	// Only visit links on the sidebar
	c.OnHTML("#menu-calorie-tables", func(e *colly.HTMLElement) {
		// Finding valid links to visit
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			link := (el.ChildAttr("a", "href"))

			getFoodInfo.Visit(link)

		})

	})

	// Parse food and its calorie into struct then json then save the file
	getFoodInfo.OnHTML(`body`, func(e *colly.HTMLElement) {
		// fmt.Println(e.Text)
		tmpFoodList := FoodList{}
		foodType := (e.ChildText(".page-title"))

		tmpFoodList.FoodType = foodType

		// Loop through each table cell
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			tmpFood := Food{}

			tmpFood.Name = el.ChildText("td:first-child")
			tmpFood.Calorie = el.ChildText("td:nth-child(5)")

			tmpFoodList.Foods = append(tmpFoodList.Foods, tmpFood)

		})

		// Converts to JSON
		js, err := json.MarshalIndent(tmpFoodList, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(js))

		// Write file to HD
		_ = ioutil.WriteFile(foodType+".json", js, 0644)

	})

	// Printing the link which colly is visiting
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Make the scrapping async
	c.Wait()

	// Start scraping
	c.Visit("https://www.calories.info/")

}
