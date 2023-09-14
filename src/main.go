package main

import (
	"fmt"

	parsing "github.com/zezoamr/golang-searchengine/parsing"
)

func main() {
	crawlLinks := []string{"https://en.wikipedia.org/wiki/Supersampling", "https://www.google.com"}
	links, words, _ := parsing.Crawl(crawlLinks)
	fmt.Println(links)
	fmt.Println(" ")
	fmt.Println(words)
}
