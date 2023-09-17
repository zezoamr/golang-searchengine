package parsing

import "fmt"

// crawlPage crawls a web page and sends the parsed page information to the parsedPagesChannel.
//
// url: The URL of the page to be crawled.
// parsedPagesChannel: A channel to send the parsed page information to.
// The function does not return anything.
func crawlPage(url string, parsedPagesChannel chan Page) {
	if url == "" {
		parsedPagesChannel <- Page{Url: url, Links: []string{}, Words: "", Err: fmt.Errorf("url is empty")}
		return
	}

	page, err := getPage(url)
	if err != nil {
		parsedPagesChannel <- Page{Url: url, Links: []string{}, Words: "", Err: err}
		return
	}

	tempLinks, tempWords, err := parsePage(url, page)
	if err != nil {
		parsedPagesChannel <- Page{Url: url, Links: []string{}, Words: "", Err: err}
		return
	}

	parsedPagesChannel <- Page{Url: url, Links: tempLinks, Words: tempWords, Err: err}
}
