package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var dogGroups []DogGroup

func homeHandler(w http.ResponseWriter, r *http.Request) {
	dogBreedsComponent := groupsListComponent(dogGroups)
	page(dogBreedsComponent).Render(r.Context(), w)
}

func groupHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["groupID"]
	params := r.URL.Query()
	groupName := params.Get("groupName")

	pageNum := 1
	pageSize := 10
	if params.Get("pageNum") != "" {
		pageNum, _ = strconv.Atoi(params.Get("pageNum"))
	}
	if params.Get("pageSize") != "" {
		pageSize, _ = strconv.Atoi(params.Get("pageSize"))
	}

	breedsList, err := getDogBreeds(groupID)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error getting dog breeds: %s\n", err)))
	}
	dogBreedsComponent := breedsListComponent(groupName, groupID, breedsList, pageNum, pageSize)
	page(dogBreedsComponent).Render(r.Context(), w)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {}

func main() {

	// Register handlers
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/favicon.ico", faviconHandler)
	r.HandleFunc("/group/{groupID}", groupHandler)

	// Practice golang HTTP and JSON unmarshalling
	var err error
	dogGroups, err = getDogGroups()
	if err != nil {
		log.Panic(fmt.Printf("Shutting down server:%s\n", err))
	}

	// Listen and serve
	log.Fatal(http.ListenAndServe(":8080", r))
}
