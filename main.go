package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// "encoding/json"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
var parks []Park

func getParks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parks)
}

func getPark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range parks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Park{})
}

func createPark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var park Park
	_ = json.NewDecoder(r.Body).Decode(&park)
	park.ID = strconv.Itoa(rand.Intn(10000000)) // Pas safe en production -
	parks = append(parks, park)
	json.NewEncoder(w).Encode(park)
}

func updatePark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range parks {
		if item.ID == params["id"] {
			parks = append(parks[:index], parks[index+1:]...)
			var park Park
			_ = json.NewDecoder(r.Body).Decode(&park)
			park.ID = params["id"]
			parks = append(parks, park)
			json.NewEncoder(w).Encode(park)
			break
		}
	}

	json.NewEncoder(w).Encode(parks)
}

func deletePark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range parks {
		if item.ID == params["id"] {
			parks = append(parks[:index], parks[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(parks)
}

// Park Struct (Model)
type Park struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	InPark       string `json:"inpark"`
	Place        string `json:"place"`
	Manufacturer string `json:"manufacturer"`
}

func main() {

	r := mux.NewRouter()

	parks = append(parks, Park{ID: "1", Name: "Osiri", InPark: "Asterix", Place: "France", Manufacturer: "Vortex"})
	parks = append(parks, Park{ID: "2", Name: "TheMonster", InPark: "Walygator Parc", Place: "France", Manufacturer: "Vortex"})
	parks = append(parks, Park{ID: "3", Name: "Taron", InPark: "Phantasialand", Place: "Allemagne", Manufacturer: "Vortex"})

	r.HandleFunc("/api/books", getParks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getPark).Methods("GET")
	r.HandleFunc("/api/books/", createPark).Methods("POST")
	r.HandleFunc("/api/books/{id}", updatePark).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deletePark).Methods("DELETE")

	// Permet de logguer l'erreur
	log.Fatal(http.ListenAndServe(":8000", r))
}
