package api

import (
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/zezoamr/golang-searchengine/searchengine"
)

func searchRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", handleSearch)
	r.Post("/createIndex/{index}", handleCreateIndex)
	r.Delete("/deleteIndex/{index}", handleDeleteIndex)
	return r
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	_, client, err := searchengine.GetClient()
	if err != nil {
		return
	}
	resultsCountStr := r.URL.Query().Get("resultsCount")
	resultsCount, err := strconv.Atoi(resultsCountStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	searchTermStr := r.URL.Query().Get("searchTerm")

	resp, err := searchengine.Search(client, "search", searchTermStr, resultsCount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(body)
}

func handleCreateIndex(w http.ResponseWriter, r *http.Request) {
	indexStr := chi.URLParam(r, "index")
	typedclient, _, err := searchengine.GetClient()
	if err != nil {
		return
	}
	err = searchengine.CreateIndex(typedclient, indexStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("index created"))
}

func handleDeleteIndex(w http.ResponseWriter, r *http.Request) {
	indexStr := chi.URLParam(r, "index")
	typedclient, _, err := searchengine.GetClient()
	if err != nil {
		return
	}
	err = searchengine.DeleteIndex(typedclient, indexStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("index deleted"))
}
