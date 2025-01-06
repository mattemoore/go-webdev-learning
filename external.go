package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

func getDogGroups() ([]DogGroup, error) {
	dogGroups := []DogGroup{}

	resp, err := http.Get("https://dogapi.dog/api/v2/groups")
	if err != nil {
		log.Println(fmt.Println(err))
		return dogGroups, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(fmt.Println(err))
		return dogGroups, err
	}

	var data DogGroupsApiData
	json.Unmarshal(body, &data)
	for _, group := range data.Data {
		dogGroups = append(dogGroups, group)
	}
	fmt.Printf("Dog groups: %v\n", dogGroups[0])
	return dogGroups, nil
}

func getDogBreeds() ([]DogBreedAttributes, error) {
	dogBreeds := []DogBreedAttributes{}

	// Request list of dog breeds
	resp, err := http.Get("https://dogapi.dog/api/v2/breeds")
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
	var data DogBreedsApiData
	json.Unmarshal(body, &data)
	for _, breed := range data.Data {
		dogBreeds = append(dogBreeds, breed.Attributes)
	}
	sort.Sort(BreedsByName(dogBreeds))
	fmt.Printf("Dog breeds: %v\n", dogBreeds[0])
	return dogBreeds, nil
}
