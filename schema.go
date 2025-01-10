package main

type DogGroupsApiData struct {
	Data []DogGroup `json:"data"`
}

type DogGroupApiData struct {
	Data DogGroup `json:"data"`
}

type DogBreeedApiData struct {
	Data DogBreed `json:"data"`
}

type DogGroup struct {
	ID            string                `json:"id"`
	Attributes    DogGroupAttributes    `json:"attributes"`
	Relationships DogGroupRelationships `json:"relationships"`
}

type DogGroupRelationships struct {
	BreedRelationships DogGroupBreedRelationships `json:"breeds"`
}

type DogGroupBreedRelationships struct {
	DogBreeds []DogBreed `json:"data"`
}

type DogGroupAttributes struct {
	Name string `json:"name"`
}

type DogBreed struct {
	ID         string             `json:"id"`
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

type GroupsByName []DogGroup

func (groups GroupsByName) Len() int { return len(groups) }
func (groups GroupsByName) Less(i, j int) bool {
	return groups[i].Attributes.Name < groups[j].Attributes.Name
}
func (groups GroupsByName) Swap(i, j int) { groups[i], groups[j] = groups[j], groups[i] }

type BreedsByName []DogBreed

func (breeds BreedsByName) Len() int { return len(breeds) }
func (breeds BreedsByName) Less(i, j int) bool {
	return breeds[i].Attributes.Name < breeds[j].Attributes.Name
}
func (breeds BreedsByName) Swap(i, j int) { breeds[i], breeds[j] = breeds[j], breeds[i] }
