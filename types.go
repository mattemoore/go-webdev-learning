package main

//#region DogGroup types

type DogGroupsApiData struct {
	Data []DogGroup `json:"data"`
}

type DogGroupApiData struct {
	Data DogGroup `json:"data"`
}

type DogGroup struct {
	ID            string                       `json:"id"`
	Attributes    DogGroupAttributes           `json:"attributes"`
	Relationships DogBreedRelationshipsInGroup `json:"relationships"`
}

type DogGroupAttributes struct {
	Name string `json:"name"`
}

type DogBreedRelationshipsInGroup struct {
	BreedsHolder DogBreedsInGroup `json:"breeds"`
}

type DogBreedsInGroup struct {
	Breeds []DogBreed `json:"data"`
}

//#endregion

//#region DogBreed types

type DogBreeedApiData struct {
	Data DogBreed `json:"data"`
}

type DogBreed struct {
	ID            string                        `json:"id"`
	Attributes    DogBreedAttributes            `json:"attributes"`
	Relationships DogGroupRelationshipsForBreed `json:"relationships"`
}

type DogBreedAttributes struct {
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Life           DogBreedMaxMin `json:"life"`
	MaleWeight     DogBreedMaxMin `json:"male_weight"`
	FemaleWeight   DogBreedMaxMin `json:"female_weight"`
	HypoAllergenic bool           `json:"hypoallergenic"`
}

type DogGroupRelationshipsForBreed struct {
	Relationships DogGroup `json:"group"`
}

type DogGroupForBreed struct {
	Group DogGroup `json:"data"`
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

//#endregion
