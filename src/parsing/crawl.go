package parsing

func crawl(urls []string) ([]string, []string, error) {
	var links []string
	var words []string
	var err error
	for _, url := range urls {
		page, err := getPage(url)
		if err != nil {
			return links, words, err
		}
		tempLinks, tempWords, err := parsePage(url, page)
		if err != nil {
			return links, words, err
		}
		links = append(links, tempLinks...)
		words = append(words, tempWords...)
	}
	return links, words, err
}
