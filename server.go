package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type DogBreed struct {
	Name      string
	SubBreeds []string
}

type BreedsByName []DogBreed

func (db BreedsByName) Len() int           { return len(db) }
func (db BreedsByName) Less(i, j int) bool { return db[i].Name < db[j].Name }
func (db BreedsByName) Swap(i, j int)      { db[i], db[j] = db[j], db[i] }

var dogBreeds BreedsByName

func homeHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	pageNum := 1
	pageSize := 10
	if params.Get("pageNum") != "" {
		pageNum, _ = strconv.Atoi(params.Get("pageNum"))
	}
	if params.Get("pageSize") != "" {
		pageSize, _ = strconv.Atoi(params.Get("pageSize"))
	}
	dogBreedsComponent := breedsListComponent(dogBreeds, pageNum, pageSize)
	page(dogBreedsComponent).Render(r.Context(), w)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {}

func main() {

	// Register handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)

	// Practice golang HTTP and JSON unmarshalling
	var err error
	dogBreeds, err = getDogBreeds()
	if err != nil {
		log.Panic(fmt.Printf("Shutting down server:%s\n", err))
	}

	// Listen and serve
	log.Fatal(http.ListenAndServe(":8080", nil))
}
