package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		url := r.URL.Query().Get("url")
		res, _ := http.Get(url)
		data := new(Data)

		err := m.Metabolize(res.Body, data)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(data)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
