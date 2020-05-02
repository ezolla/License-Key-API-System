package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" // external package, HTTP router for Go web servers
)

// License Struct
type License struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	User *User  `json:"user"`
}

// User Struct
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Slice of Licenses
var licenses []License

func main() {
	// Initializing Router
	router := mux.NewRouter()

	// Mock Data
	licenses = append(licenses, License{ID: "1", Key: "1111-1111-1111-1111", User: &User{Firstname: "Bart", Lastname: "Simpson"}})
	licenses = append(licenses, License{ID: "2", Key: "2222-2222-2222-2222", User: &User{Firstname: "Cole", Lastname: "Bennett"}})
	licenses = append(licenses, License{ID: "3", Key: "3333-3333-3333-3333", User: &User{Firstname: "Homer", Lastname: "Simpson"}})

	// Router Handlers "Endpoints"
	router.HandleFunc("/api/licenses", getLicenses).Methods("GET")
	router.HandleFunc("/api/licenses/{id}", searchLicense).Methods("GET")
	router.HandleFunc("/api/licenses", createLicense).Methods("POST")
	router.HandleFunc("/api/licenses/{id}", updateLicense).Methods("PUT")
	router.HandleFunc("/api/licenses/{id}", removeLicense).Methods("DELETE")

	// Run Server
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Retrieves Licenses
func getLicenses(w http.ResponseWriter, r *http.Request) {
	// Making Response in JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(licenses)
}

// Search License
func searchLicense(w http.ResponseWriter, r *http.Request) {
	// Making Response in JSON
	w.Header().Set("Content-Type", "application/json")

	// Saving Parameters
	params := mux.Vars(r)

	// Loop through IDs to Find Associated Key
	for _, item := range licenses {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&License{})
}

// Creates License
func createLicense(w http.ResponseWriter, r *http.Request) {
	// Making Response in JSON
	w.Header().Set("Content-Type", "application/json")

	var license License
	_ = json.NewDecoder(r.Body).Decode(&license)

	// Creating Random New ID
	license.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe

	// Adding New Book to List of Books
	licenses = append(licenses, license)

	// Send New Book Upon Request
	json.NewEncoder(w).Encode(license)
}

// Update License
func updateLicense(w http.ResponseWriter, r *http.Request) {
	// Making Response in JSON
	w.Header().Set("Content-Type", "application/json")

	// Saving Parameters
	params := mux.Vars(r)

	// Looping through Licenses to find ID
	for index, item := range licenses {
		if item.ID == params["id"] {
			licenses = append(licenses[:index], licenses[index+1:]...)
			var license License
			_ = json.NewDecoder(r.Body).Decode(&license)

			// Creating Random New ID
			license.ID = params["id"]

			// Adding New Book to List of Books
			licenses = append(licenses, license)

			// Send New Book Upon Request
			json.NewEncoder(w).Encode(license)

			return
		}
	}
	json.NewEncoder(w).Encode(licenses)
}

// Remove License
func removeLicense(w http.ResponseWriter, r *http.Request) {
	// Making Response in JSON
	w.Header().Set("Content-Type", "application/json")

	// Saving Parameters
	params := mux.Vars(r)

	// Finding
	for index, item := range licenses {
		if item.ID == params["id"] {
			licenses = append(licenses[:index], licenses[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(licenses)
}
