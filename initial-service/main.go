package main

import (
	"fmt"

	"github.com/jos-monotako/feedling/initial-service/scraper"
)

func main() {
	feed, _ := scraper.GetFeed("https://steamcommunity.com/games/2622380/rss/")

	for _, item := range feed.Channel.Items {
		fmt.Println(item.Title)
		fmt.Println(scraper.HTMLtoCleanText(item.Description))
		fmt.Println("===================================================")
	}

	//server.Start()
}
