package main

type DogGroupsApiData struct {
	Data []DogGroup `json:"data"`
}

type DogGroup struct {
	ID            string                `json:"id"`
	Type          string                `json:"type"`
	Attributes    DogGroupAttributes    `json:"attributes"`
	Relationships DogGroupRelationships `json:"relationships"`
}

type DogGroupRelationships struct {
	DogGroupBreeds DogGroupBreeds `json:"breeds"`
}

type DogGroupBreeds struct {
	DogGroupBreedsData []DogGroupBreed `json:"data"`
}

type DogGroupBreed struct {
	ID string `json:"id"`
}

type DogGroupAttributes struct {
	Name string `json:"name"`
}

type DogBreedsApiData struct {
	Data []DogBreed `json:"data"`
}

type DogBreed struct {
	ID         string             `json:"id"`
	Type       string             `json:"type"`
	Attributes DogBreedAttributes `json:"attributes"`
	// Relationships DogBreedRelationships `json:"relationships"`
}

type DogBreedAttributes struct {
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Life           DogBreedMaxMin `json:"life"`
	MaleWeight     DogBreedMaxMin `json:"male_weight"`
	FemaleWeight   DogBreedMaxMin `json:"female_weight"`
	HypoAllergenic bool           `json:"hypoallergenic"`
}

type DogBreedMaxMin struct {
	Max int `json:"max"`
	Min int `json:"min"`
}

type BreedsByName []DogBreedAttributes

func (db BreedsByName) Len() int           { return len(db) }
func (db BreedsByName) Less(i, j int) bool { return db[i].Name < db[j].Name }
func (db BreedsByName) Swap(i, j int)      { db[i], db[j] = db[j], db[i] }
