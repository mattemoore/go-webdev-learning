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

	cachedDogGroups, err := fetchDogGroups()

	if err == nil && len(cachedDogGroups) > 0 {
		return cachedDogGroups, nil
	}

	log.Println("Cannot retrieve dog groups from cache.  Calling API...")

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

	cacheDogGroups(dogGroups)

	return dogGroups, nil
}

type DogGroupRedis struct {
	ID   string `redis:"id"`
	Name string `redis:"name"`
}

func fetchDogGroups() ([]DogGroup, error) {
	var cachedDogGroups []DogGroupRedis
	var dogGroups []DogGroup

	iter := redisClient.Scan(ctx, 0, "dog:group:*", 0).Iterator()
	for iter.Next(ctx) {
		log.Println("key: ", iter.Val())
		key := iter.Val()
		var dogGroupRedis DogGroupRedis
		error := redisClient.HGetAll(ctx, key).Scan(&dogGroupRedis)
		if error != nil {
			return nil, error
		}
		dogGroup := DogGroup{
			ID:         dogGroupRedis.ID,
			Attributes: DogGroupAttributes{Name: dogGroupRedis.Name},
		}
		dogGroups = append(dogGroups, dogGroup)

	}
	if err := iter.Err(); err != nil {
		log.Printf("Dog groups retrieved from cache: %v\n", cachedDogGroups)
		return nil, err
	}

	log.Println("Cache hit for dog groups.")
	return dogGroups, nil
}

func cacheDogGroups(dogGroups []DogGroup) {
	log.Println("Cache miss for dog groups.")
	var err error = nil
	for _, group := range dogGroups {
		err = redisClient.HSet(ctx, fmt.Sprintf("dog:group:%s", group.ID), "name", group.Attributes.Name, "id", group.ID).Err()
	}

	if err != nil {
		log.Printf("Error caching dog groups: %s\n", err)
	}
}

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
