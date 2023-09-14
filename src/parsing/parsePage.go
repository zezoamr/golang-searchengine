package parsing

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// parsePage takes a string as input and returns two slices of strings and an error.
//
// The function parses the given string as HTML and extracts all the links and words
// from it. It returns the extracted links and words as slices of strings. If there
// is an error in parsing the HTML, it returns an empty slice for both links and words
// and the error.
func parsePage(originalURL string, body string) ([]string, []string, error) {
	parsedLinks := []string{}
	parsedWords := []string{}
	doc, err := html.Parse(strings.NewReader((body)))
	if err != nil {
		return []string{}, []string{}, err
	}

	baseURL, err := url.Parse(originalURL)
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
				if attr.Key == "href" {
					resolvedURL, err := baseURL.Parse(attr.Val)
					if err != nil {
						continue // skip this link
					}
					parsedLinks = append(parsedLinks, resolvedURL.String())
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
