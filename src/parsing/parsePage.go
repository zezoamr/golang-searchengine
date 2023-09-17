package parsing

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// parsePage parses the given page and extracts the links and words from it.
//
// Parameters:
// - originalURL: the original URL of the page to parse (string).
// - body: the body of the page to parse (string).
//
// Returns:
// - parsedLinks: a slice of strings containing the parsed links from the page ([]string).
// - parsedWords: a string containing the parsed words from the page.
// - error: an error object if any error occurs during the parsing process.
func parsePage(originalURL string, body string) ([]string, string, error) {
	parsedLinks := []string{}
	var parsedWords strings.Builder
	doc, err := html.Parse(strings.NewReader((body)))
	if err != nil {
		return []string{}, "", err
	}

	baseURL, err := url.Parse(originalURL)
	if err != nil {
		return []string{}, "", err
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
			if strings.TrimSpace(node.Data) != "" {
				parsedWords.WriteString(cleanText(strings.TrimSpace(node.Data)))
				parsedWords.WriteString(" ")
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return parsedLinks, strings.TrimSpace(parsedWords.String()), nil
}
