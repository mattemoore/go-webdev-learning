package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var dogGroups []DogGroup

func homeHandler(w http.ResponseWriter, r *http.Request) {
	dogBreedsComponent := groupsListComponent(dogGroups)
	page(dogBreedsComponent).Render(r.Context(), w)
}

func groupHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	groupID := params.Get("groupID")
	w.Write([]byte(fmt.Sprintf("Group ID: %s\n", groupID)))

	pageNum := 1
	pageSize := 10
	if params.Get("pageNum") != "" {
		pageNum, _ = strconv.Atoi(params.Get("pageNum"))
	}
	if params.Get("pageSize") != "" {
		pageSize, _ = strconv.Atoi(params.Get("pageSize"))
	}
	w.Write([]byte(fmt.Sprintf("Page number: %d\n", pageNum)))
	w.Write([]byte(fmt.Sprintf("Page size: %d\n", pageSize)))

	// TODO: Query api to get list of breeds for the breedID
	// dogBreedsComponent := breedsListComponent(breedsList, pageNum, pageSize)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {}

func main() {

	// Register handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/group/", groupHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)

	// Practice golang HTTP and JSON unmarshalling
	var err error
	dogGroups, err = getDogGroups()
	if err != nil {
		log.Panic(fmt.Printf("Shutting down server:%s\n", err))
	}

	// Listen and serve
	log.Fatal(http.ListenAndServe(":8080", nil))
}
