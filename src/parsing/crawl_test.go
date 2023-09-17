package parsing

//unittests not done yet

// func TestCrawl(t *testing.T) {
// 	//testcase 1: max total pages
// 	crawlLinks := []string{"https://en.wikipedia.org/wiki/Supersampling"}
// 	p := crawl(crawlLinks, 4, 2, 5, 5)
// 	if len(p) > 0 {
// 		fmt.Println(len(p))
// 		//fmt.Println(p[0].Links)
// 	} else {
// 		fmt.Println("No pages were crawled.")
// 	}

// 	//testcase 2: max per page

// }

// func TestCrawl2(t *testing.T) {

// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		html := "<html><body><p>Paragraph 1</p><p>Paragraph 2</p></body></html>"
// 		w.Write([]byte(html))
// 	}))
// 	defer server.Close()
// 	crawlLinks := []string{server.URL}
// 	p := crawl(crawlLinks, 4, 2, 5, 5)
// 	if len(p) != 1 {
// 		t.Errorf("No pages were crawled.")
// 	}
// }

// func TestCrawl3(t *testing.T) {
// 	urls := []string{"https://example.com", "https://google.com"}

// 	t.Run("Number of URLs must be between 1 and 100", func(t *testing.T) {
// 		invalidURLs := []string{}
// 		for i := 0; i < 101; i++ {
// 			invalidURLs = append(invalidURLs, "https://example.com")
// 		}

// 		pages := crawl(invalidURLs, 10, 10, 1, 1)

// 		if len(pages) != 0 {
// 			t.Errorf("Expected 0 pages, got %d", len(pages))
// 		}
// 	})

// 	t.Run("Crawl pages and process links", func(t *testing.T) {
// 		pages := crawl(urls, 3, 2, 1, 1)

// 		if len(pages) != 3 {
// 			t.Errorf("Expected 3 pages, got %d", len(pages))
// 		}

// 		// Assert that each page has at most 2 links processed
// 		for _, page := range pages {
// 			if len(page.Links) > 2 {
// 				t.Errorf("Expected at most 2 links, got %d", len(page.Links))
// 			}
// 		}
// 	})

// 	t.Run("No more links to process", func(t *testing.T) {
// 		start := time.Now()
// 		pages := crawl(urls, 10, 10, 1, 1)
// 		elapsed := time.Since(start)

// 		if len(pages) != 10 {
// 			t.Errorf("Expected 10 pages, got %d", len(pages))
// 		}

// 		if elapsed < 10*time.Second {
// 			t.Errorf("Expected crawl to take at least 10 seconds, took %s", elapsed)
// 		}
// 	})
// }
