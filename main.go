package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


var indexTemplate = template.Must(template.ParseFiles("index.tmpl"))

// Index is a data structure used to populate an indexTemplate.
type Index struct {
	Title string
	Body  string
	Links []Link
}

type Link struct {
	URL, Title string
}

// indexHandler is an HTTP handler that serves the index page.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := &Index{
		Title: "Microservice framework",
		Body:  "Welcome to the image gallery.",
		Links: []Link {
			{
				URL:   "https://servicecomb.apache.org",
				Title: "ServiceComb",
			},
				{
				URL:   "https://spring.io/projects/spring-cloud",
				Title: "Springcloud",
			},
		},
	}

	if err := indexTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}
