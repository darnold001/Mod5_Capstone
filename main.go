package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//User Struct (Model)
type Property struct {
	ID       string    `json:"id"`
	IDname   string    `json:"idname"`
	Address  string    `json:"address"`
	Username *Username `json:"username"`
}

//Init book variables as a a slice Property Struct
var properties []Property

//Username Struct
type Username struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      string `json:"role"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

//Get all Properties
func getProperties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(properties)
}

// get single property
func getProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params
	//Loop through properties ansd find ID
	for _, item := range properties {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Property{})
}

//create property
func createProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var property Property
	_ = json.NewDecoder(r.Body).Decode(&property)
	property.ID = strconv.Itoa(rand.Intn(10000)) // Mock ID - Could create duplicates
	properties = append(properties, property)
	json.NewEncoder(w).Encode(property)

}

//update Property
func updateProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range properties {
		if item.ID == params["id"] {
			properties = append(properties[:index], properties[index+1:]...)
			var property Property
			_ = json.NewDecoder(r.Body).Decode(&property)
			property.ID = params["id"]
			properties = append(properties, property)
			json.NewEncoder(w).Encode(property)
			return
		}
	}
	json.NewEncoder(w).Encode(properties)

}
func deleteProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range properties {
		if item.ID == params["id"] {
			properties = append(properties[:index], properties[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(properties)
}

func main() {
	//Init Router
	router := mux.NewRouter()

	// Mock Data
	properties = append(properties, Property{ID: "1", IDname: "LoHi House", Address: "1834 Platte St. Denver, CO", Username: &Username{ID: 1, Firstname: "Devin", Lastname: "Arnold", Role: "Tenant", Phone: "(303) 519-9508", Email: "dbarnold432@gmail.com"}})
	properties = append(properties, Property{ID: "2", IDname: "Barnum House", Address: "184 Utica St. Denver, CO", Username: &Username{ID: 2, Firstname: "Kersten", Lastname: "Arnold", Role: "Tenant", Phone: "(303) 419-9508", Email: "kersnt@gmail.com"}})
	properties = append(properties, Property{ID: "3", IDname: "DU House", Address: "1834 Platte St. Denver, CO", Username: &Username{ID: 3, Firstname: "Colin", Lastname: "Arnold", Role: "Landlord", Phone: "(303) 545-9508", Email: "COA@gmail.com"}})

	//Route Handlers
	router.HandleFunc("/api/properties", getProperties).Methods("GET")
	router.HandleFunc("/api/properties/{id}", getProperty).Methods("GET")
	router.HandleFunc("/api/properties", createProperty).Methods("POST")
	router.HandleFunc("/api/properties/{id}", updateProperty).Methods("PUT")
	router.HandleFunc("/api/properties/{id}", deleteProperty).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
