package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

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
	fmt.Printf("Dog breeds: %v\n", dogBreeds)
	return dogBreeds, nil
}
