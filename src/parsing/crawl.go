package parsing

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

type Page struct {
	Url   string   `json:"url"`
	Links []string `json:"links"`
	Words string   `json:"words"`
	Err   error    `json:"-"`
}

func crawl(urls []string, MAX_PAGES_TO_BE_PARSED int, MAX_LINKS_PER_PAGE int, channelsLength int, rateLimit int) []Page {

	if len(urls) == 0 || len(urls) > channelsLength {
		log.Fatal("number of starting URLs must be between 1 and linksChannelLength")
	}

	parsedChannel := make(chan Page, channelsLength)
	linksChannel := make(chan string, channelsLength)
	done := make(chan bool)
	parsedPagesCount := 0
	PAGES := []Page{}
	crawledPages := make(map[string]struct{})
	mutex := &sync.RWMutex{}
	activeRequests := int64(0)

	for _, url := range urls {
		normalizedurl, err := normalizeURL(url)
		if err != nil {
			log.Println("Error normalizing URL:", err)
			continue
		}
		linksChannel <- normalizedurl
	}

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
				activeRequests--
				mutex.Unlock()

				log.Println("finished crawling a page", tempPage.Url)

				if parsedPagesCount > MAX_PAGES_TO_BE_PARSED {
					done <- true
					return
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
					mutex.RLock()
					_, ok := crawledPages[link]
					mutex.RUnlock()
					if !ok {
						count++
						linksChannel <- normalizedLink
					}
				}
			}(tempPage)
		}
	}()

	sem := semaphore.NewWeighted(int64(rateLimit))

	go func() {
		for link := range linksChannel {
			if err := sem.Acquire(context.Background(), 1); err != nil {
				log.Printf("Failed to acquire semaphore: %v", err)
				continue
			}
			go func(link string) {
				defer sem.Release(1)
				fmt.Println("crawling a new link", link)
				activeRequests++
				crawlPage(link, parsedChannel)
			}(link)
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			if len(linksChannel) == 0 && activeRequests == 0 {
				log.Println("No more links to process")
				done <- true
				return
			}
			// Check if there are no more links to process and no more active requests. to avoid when there is no more links to be parsed while still lessen than MAX_PAGES_TO_BE_PARSED
		}
	}()

	<-done

	close(linksChannel)
	close(parsedChannel)
	return PAGES
}
