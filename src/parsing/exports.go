package parsing

var LINKS_CHANNEL_LENGTH = 100
var PARSED_PAGES_CHANNEL_LENGTH = 100

func Crawl(urls []string, MAX_PAGES_TO_BE_PARSED int, MAX_LINKS_PER_PAGE int, lc int, pc int) []Page {
	return crawl(urls, MAX_PAGES_TO_BE_PARSED, MAX_LINKS_PER_PAGE, lc, pc) //, LINKS_CHANNEL_LENGTH, PARSED_PAGES_CHANNEL_LENGTH)
}
