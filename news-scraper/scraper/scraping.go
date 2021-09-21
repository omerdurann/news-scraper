package scraper

import (
	"fmt"
	"net/http"
	"time"
	"news-scraper/helpers"

	"github.com/mmcdole/gofeed"
)

func StartRSSCrawler(agencies []helpers.Agancy) {
	// 0 olacak demo için len kullandım
	// index := len(resources) - 1
	index := 0
	for {
		if index >= len(agencies) {
			time.Sleep(1 * time.Hour)
			index = 0
		}

		source := agencies[index]
		go _handle_rss_list(source)
		index += 1
		time.Sleep(300 * time.Millisecond)
	}
}

// Her bir RSS kategorisi için loop
func _handle_rss_list(source helpers.Agancy) {
	index := 0

	for {
		if index >= len(source.RSS) {
			fmt.Println("/////////////////////////////////////////////////////////////")
			fmt.Println(source.Agency + " için rss ler bitti. 1. duraklama başladı!")
			fmt.Println("/////////////////////////////////////////////////////////////")
			time.Sleep(1 * time.Hour)
			index = 0
		}

		source := source.RSS[index]
		feed, err := _get_rss_list_by_category(source)
		if err != nil {
			fmt.Println("Hata!!!!! : ", err)
			index += 1
			time.Sleep(30 * time.Minute)
		}
		for k := 0; k < len(feed); k++ {
			feed_handler(feed[k])
			time.Sleep(1 * time.Millisecond)
		}
		index += 1
	}
}

// BU KATEGORİDEKİ RSS LİNKİNE GİDİP HER ŞEYİ SCRAPER EDİP---
// SONRAKİ KATEGORİYE GEÇMEK İÇİN RETURN EDECEK...
func _get_rss_list_by_category(rss helpers.RssOBJ) ([]*gofeed.Item, error) {
	fp := gofeed.NewParser()
	fp.Client = &http.Client{Timeout: time.Second * 10}

	feed, err := fp.ParseURL(rss.Source)
	if err != nil {
		return nil, err
	}

	return feed.Items, nil
}

// her bir rss haberini işlememize yarar
func feed_handler(news *gofeed.Item) {
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("News Categories= ", news.Categories)
	fmt.Println("News Title= ", news.Title)
	fmt.Println("News Description= ", news.Description)
	fmt.Println("News Link= ", news.Link)
	fmt.Println("-------------------------------------------------------------")
}
