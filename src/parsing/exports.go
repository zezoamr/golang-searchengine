package parsing

func Crawl(urls []string, MAX_PAGES_TO_BE_PARSED int, MAX_LINKS_PER_PAGE int, channelsLength int, rateLimit int) []Page {
	return crawl(urls, MAX_PAGES_TO_BE_PARSED, MAX_LINKS_PER_PAGE, channelsLength, rateLimit)
}
