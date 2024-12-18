package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
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
	dogBreedsComponent := breedsListComponent(dogBreeds)
	page(dogBreedsComponent).Render(r.Context(), w)
}

func main() {

	// Register handlers
	http.HandleFunc("/", homeHandler)

	// Practice golang HTTP and JSON unmarshalling
	var err error
	dogBreeds, err = getDogBreeds()
	if err != nil {
		log.Panic(fmt.Printf("Shutting down server:%s\n", err))
	}

	// Listen and serve
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getDogBreeds() ([]DogBreed, error) {
	dogBreeds := []DogBreed{}

	// Request list of dog breeds
	resp, err := http.Get("https://dog.ceo/api/breeds/list/all")
	if err != nil {
		log.Println(fmt.Println(err))
		return dogBreeds, err
	}

	// Read response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(fmt.Println(err))
		return dogBreeds, err
	}

	// Convert JSON response into map with breeds as keys and sub breeds as values
	var result map[string]any
	json.Unmarshal(body, &result)
	message := result["message"].(map[string]any)
	for key, value := range message {
		subBreeds := make([]string, 0)
		for _, v := range value.([]any) {
			subBreeds = append(subBreeds, v.(string))
		}
		dogBreeds = append(dogBreeds, DogBreed{Name: key, SubBreeds: subBreeds})
	}

	sort.Sort(BreedsByName(dogBreeds))
	return dogBreeds, nil
}
