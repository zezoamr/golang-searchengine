package parsing

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrawl(t *testing.T) {
	//testcase 1: max total pages
	crawlLinks := []string{"https://en.wikipedia.org/wiki/Supersampling"}
	p := crawl(crawlLinks, 4, 2, 5, 5)
	if len(p) != 4 {
		t.Log(len(p))
	} else {
		t.Logf("incorrect number of pages: %d", len(p))
	}

	//testcase 2: no deadlock when crawling multiple pages where available pages is less than max total pages requested

}

func TestCrawl2(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		html := "<html><body><p>Paragraph 1</p><p>Paragraph 2</p></body></html>"
		w.Write([]byte(html))
	}))
	defer server.Close()
	crawlLinks := []string{server.URL}
	p := crawl(crawlLinks, 4, 2, 5, 5)
	if len(p) != 1 {
		t.Errorf("No pages were crawled.")
	}
}
