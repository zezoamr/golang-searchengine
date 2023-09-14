package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	crawlLinks := []string{"https://en.wikipedia.org/wiki/Supersampling", "https://www.google.com"}
	links, words, _ := crawl(crawlLinks)
	fmt.Println(links)
	fmt.Println(" ")
	fmt.Println(words)
}

// getPage retrieves the HTML content from the specified URL.
//
// It takes a single parameter:
// - url: a string representing the URL to fetch the HTML from.
//
// The function returns a string containing the HTML content.
func getPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()

	// following reads all at once
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	formattedString := string(body)
	return formattedString, nil

	// following reads in chunks
	// buf := make([]byte, 1024)
	// for {
	// 	n, err := resp.Body.Read(buf)
	// 	if err != nil && err != io.EOF {
	// 		log.Fatalln("Error:", err)
	// 	}
	// 	if n == 0 {
	// 		break
	// 	}
	// 	fmt.Print(string(buf[:n]))
	// }
}

// parsePage takes a string as input and returns two slices of strings and an error.
//
// The function parses the given string as HTML and extracts all the links and words
// from it. It returns the extracted links and words as slices of strings. If there
// is an error in parsing the HTML, it returns an empty slice for both links and words
// and the error.
func parsePage(body string) ([]string, []string, error) {
	parsedLinks := []string{}
	parsedWords := []string{}
	doc, err := html.Parse(strings.NewReader((body)))
	if err != nil {
		return []string{}, []string{}, err
	}

	filterTags := map[string]bool{ // extract the variable outside so it doesn't get created every time
		"a":        true,
		"script":   true,
		"style":    true,
		"noscript": true,
		"iframe":   true,
		"object":   true,
		"embed":    true,
		"param":    true,
	}

	var f func(*html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" { //&& (strings.HasPrefix(attr.Val, "http") || strings.HasPrefix(attr.Val, "https")) {
					parsedLinks = append(parsedLinks, attr.Val)
				}
			}
		}
		if node.Type == html.TextNode && (node.Parent == nil || !filterTags[node.Parent.Data]) {
			parsedWords = append(parsedWords, strings.TrimSpace(node.Data))
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return parsedLinks, parsedWords, nil
}

func crawl(urls []string) ([]string, []string, error) {
	var links []string
	var words []string
	var err error
	for _, url := range urls {
		page, err := getPage(url)
		if err != nil {
			return links, words, err
		}
		tempLinks, tempWords, err := parsePage(page)
		if err != nil {
			return links, words, err
		}
		links = append(links, tempLinks...)
		words = append(words, tempWords...)
	}
	return links, words, err
}
