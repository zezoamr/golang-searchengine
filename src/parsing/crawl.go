package parsing

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Page struct {
	Url   string
	Links []string
	Words string
	Err   error
}

func crawl(urls []string, MAX_PAGES_TO_BE_PARSED int, MAX_LINKS_PER_PAGE int, linksChannelLength int, parsedPagesChannelLength int) []Page {

	if len(urls) == 0 || len(urls) > linksChannelLength {
		log.Fatal("number of starting URLs must be between 1 and linksChannelLength")
	}

	parsedChannel := make(chan Page, linksChannelLength)
	linksChannel := make(chan string, parsedPagesChannelLength)
	done := make(chan bool)
	parsedPagesCount := 0
	PAGES := []Page{}
	crawledPages := make(map[string]struct{})
	mutex := &sync.Mutex{}
	for _, url := range urls {
		normalizedurl, err := normalizeURL(url)
		if err != nil {
			log.Println("Error normalizing URL:", err)
			continue
		}
		linksChannel <- normalizedurl
	}

	go func() {
		for {
			time.Sleep(1 * time.Second)
			if len(linksChannel) == 0 {
				log.Println("No more links to process")
				done <- true
				return
			}
			// checks if there are no more links to process. If the links channel is empty, it breaks out of the for loop.
			// This is to prevent the program from getting stuck in an infinite loop if there are not enough links to reach MAX_PAGES_TO_BE_PARSED.
		}
	}()

	go func() {
		for tempPage := range parsedChannel {
			go func(tempPage Page) {

				normalizedLink, err := normalizeURL(tempPage.Url)
				if err != nil {
					normalizedLink = tempPage.Url
				}

				mutex.Lock()
				crawledPages[normalizedLink] = struct{}{}
				PAGES = append(PAGES, tempPage)
				parsedPagesCount++
				mutex.Unlock()

				log.Println("finished crawling a page", tempPage.Url)

				if parsedPagesCount > MAX_PAGES_TO_BE_PARSED {
					done <- true
				}

				count := 0
				for _, link := range tempPage.Links {
					if count == MAX_LINKS_PER_PAGE {
						break
					}
					normalizedLink, err := normalizeURL(link)
					if err != nil {
						log.Println("Error normalizing URL:", err)
						continue
					}
					mutex.Lock()
					_, ok := crawledPages[link]
					mutex.Unlock()
					if !ok {
						count++
						linksChannel <- normalizedLink
					}
				}
			}(tempPage)
		}
	}()

	go func() {
		for link := range linksChannel {
			fmt.Println("crawling a new link", link)
			go crawlPage(link, parsedChannel)
		}
	}()

	<-done

	close(linksChannel)
	close(parsedChannel)
	return PAGES
}
