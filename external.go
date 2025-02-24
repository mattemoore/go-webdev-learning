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

	v := redisClient.Get(ctx, "bike:1")
	fmt.Printf("%s\n", v)

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
	dogGroups = append(dogGroups, data.Data...)

	sort.Sort(GroupsByName(dogGroups))

	// cacheDogGroups(dogGroups)

	return dogGroups, nil
}

// func cacheDogGroups(dogGroups []DogGroup) {
// 	redis.
// }

func getDogBreeds(groupID string) ([]DogBreed, error) {
	dogBreeds := []DogBreed{}

	resp, err := http.Get(fmt.Sprintf("https://dogapi.dog/api/v2/groups/%s", groupID))
	if err != nil {
		log.Println(fmt.Println(err))
		return dogBreeds, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(fmt.Println(err))
		return dogBreeds, err
	}

	var data DogGroupApiData
	json.Unmarshal(body, &data)
	for _, breed := range data.Data.Relationships.BreedsHolder.Breeds {
		fullDogBreed, error := getDogBreed(breed.ID)
		if error != nil {
			log.Println(fmt.Println(err))
			return dogBreeds, err
		}
		dogBreeds = append(dogBreeds, fullDogBreed)
	}

	sort.Sort(BreedsByName(dogBreeds))
	return dogBreeds, nil
}

func getDogBreed(breedID string) (DogBreed, error) {
	dogBreed := DogBreed{}

	resp, err := http.Get(fmt.Sprintf("https://dogapi.dog/api/v2/breeds/%s", breedID))
	if err != nil {
		log.Println(fmt.Println(err))
		return dogBreed, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(fmt.Println(err))
		return dogBreed, err
	}

	var data DogBreeedApiData
	json.Unmarshal(body, &data)
	dogBreed = data.Data

	return dogBreed, nil
}
