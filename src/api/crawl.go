package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/zezoamr/golang-searchengine/parsing"
	"github.com/zezoamr/golang-searchengine/searchengine"
)

func crawlRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", handleCrawl)
	return r
}

func handleCrawl(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Links []string `json:"links"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	maxTotalStr := r.URL.Query().Get("maxTotal")
	maxTotal, err := strconv.Atoi(maxTotalStr) //issue here
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	maxPerPageStr := r.URL.Query().Get("maxPerPage")
	maxPerPage, err := strconv.Atoi(maxPerPageStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rateStr := r.URL.Query().Get("rate")
	rate, err := strconv.Atoi(rateStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if maxTotal > 50 {
		maxTotal = 50
	}
	if maxPerPage > 10 {
		maxPerPage = 10
	}
	if rate > 20 {
		rate = 20
	}

	go func() {
		pages := parsing.Crawl(data.Links, maxTotal, maxPerPage, 100, rate)
		typedclient, _, err := searchengine.GetClient()
		if err != nil {
			return
		}
		_, err = searchengine.CreateDocuments(typedclient, "search", pages)
		if err != nil {
			log.Println(err)
		}
	}()

	w.Write([]byte("crawling"))
}
