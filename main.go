package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	m "github.com/keighl/metabolize"
)

type Data struct {
	Title       string  `meta:"og:title" json:"title"`
	Description string  `meta:"og:description,description" json:"description"`
	Type        string  `meta:"og:type" json:"type,omitempty"`
	URL         url.URL `meta:"og:url" json:"url,omitempty"`
	Image       string  `meta:"og:image" json:"image,omitempty"`
	SiteName    string  `meta:"og:site_name" json:"site_name,omitempty"`
	Body        string  `meta:"article" json:"body"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Crawler).Methods("GET")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))

}

func Crawler(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query().Get("url")
	res, _ := http.Get(url)
	data := new(Data)

	err := m.Metabolize(res.Body, data)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(data)
}
