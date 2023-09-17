package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/zezoamr/golang-searchengine/parsing"
)

func main() {
	crawlLinks := []string{"https://en.wikipedia.org/wiki/Supersampling"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		html := "<html><body><p>Paragraph 1</p><p>Paragraph 2</p></body></html>"
		w.Write([]byte(html))
	}))
	defer server.Close()
	// crawlLinks := []string{server.URL}
	p := parsing.Crawl(crawlLinks, 4, 2, 2, 1)

	if len(p) > 0 {
		fmt.Println(len(p))
		//fmt.Println(p[0].Links)
		//fmt.Println(p[0].Words)
	} else {
		fmt.Println("No pages were crawled.")
	}
}
