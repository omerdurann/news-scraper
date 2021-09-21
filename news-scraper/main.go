package main

import (
	"fmt"
	"news-scraper/helpers"
	"news-scraper/scraper"
)

func main() {
	fmt.Println("Our program is starting...")

	agencies, err := helpers.GetRss("tr")
	if err != nil {
		fmt.Println("Hata!! :", err)
		return
	}

	//fmt.Println(agencies)
	scraper.StartRSSCrawler(agencies)
}
