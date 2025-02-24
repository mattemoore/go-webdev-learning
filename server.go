package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var ctx context.Context

func groupsPageHandler(w http.ResponseWriter, r *http.Request) {
	groupsListPlaceholder := groupsListContainer()
	page(groupsListPlaceholder).Render(r.Context(), w)
}

func groupsListHandler(w http.ResponseWriter, r *http.Request) {
	dogGroups, err := getDogGroups()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error getting dog groups: %s\n", err)))
	}

	dogBreedsComponent := groupsListComponent(dogGroups)
	dogBreedsComponent.Render(r.Context(), w)
}

func breedsPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["groupID"]
	params := r.URL.Query()
	groupName := params.Get("groupName")
	pageNum := params.Get("pageNum")
	pageSize := params.Get("pageSize")
	breedsListPlaceholder := breedsListContainer(groupName, groupID, pageNum, pageSize)
	page(breedsListPlaceholder).Render(r.Context(), w)
}

func breedsListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["groupID"]
	params := r.URL.Query()

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
	dogBreedsComponent := breedsListComponent(groupID, breedsList, pageNum, pageSize)
	dogBreedsComponent.Render(r.Context(), w)
}

// func breedHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	breedID := vars["breedID"]
// 	breed, err := getDogBreed(breedID)
// 	if err != nil {
// 		w.Write([]byte(fmt.Sprintf("Error getting dog breed: %s\n", err)))
// 	}
// 	breedComponent := breedComponent(breed)
// 	breedComponent.Render(r.Context(), w)
// }

func faviconHandler(w http.ResponseWriter, r *http.Request) {}

func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
}

func main() {

	// Create redis client
	redisClient = newRedisClient()
	ctx = context.Background()
	fmt.Printf("Redis client created: %v\n", redisClient)
	defer redisClient.Close()

	// Register handlers
	r := mux.NewRouter()
	r.HandleFunc("/", groupsPageHandler)
	r.HandleFunc("/groups", groupsListHandler)
	r.HandleFunc("/group/{groupID}", breedsPageHandler)
	r.HandleFunc("/group/list/{groupID}", breedsListHandler)
	// r.HandleFunc("/breed/{breedID}", breedHandler)
	r.HandleFunc("/favicon.ico", faviconHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Listen and serve
	log.Fatal(http.ListenAndServe(":8080", r))
}
