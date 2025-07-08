package main

import (
	"fmt"

	"github.com/jos-monotako/feedling/initial-service/scraper"
)

func main() {
	feed, _ := scraper.GetFeed("https://steamcommunity.com/games/2622380/rss/")

	fmt.Println(feed.Channel.Items[0].Description)
}
