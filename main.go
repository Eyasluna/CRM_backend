package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

// createCustomer creates a new customer
func createCustomer(ID, Name, Role, Email, Phone string, Contacted bool) (result *Customer) {
	customer := Customer{
		ID:        ID,
		Name:      Name,
		Role:      Role,
		Email:     Email,
		Phone:     Phone,
		Contacted: Contacted,
	}
	result = &customer
	return result
}

// modifyCustomer modifies an existing customer
func (c *Customer) modifyCustomer(ID, Name, Role, Email, Phone string, Contacted bool) {
	c.ID = ID
	c.Name = Name
	c.Role = Role
	c.Email = Email
	c.Phone = Phone
	c.Contacted = Contacted
}

//var customers = make(map[string]Customer)

var customers = map[string]string{
	"1": "Andy",
	"2": "Peter",
	"3": "Gabriella",
	"4": "Jordy",
}

// getCustomers returns all customers
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customers)
}

// getCustomer returns a single customer
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, ok := customers[id]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(customers)
	}
}

// addCustomer adds a new customer
func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newEntry map[string]string

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newEntry)

	for k, v := range newEntry {
		if _, ok := customers[k]; ok {
			w.WriteHeader(http.StatusConflict)
		} else {
			customers[k] = v
			w.WriteHeader(http.StatusCreated)
		}
	}

	json.NewEncoder(w).Encode(customers)
}

// updateCustomer updates a customer info
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	var newEntry map[string]string

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newEntry)

	if _, ok := customers[id]; ok {
		for k, v := range newEntry {
			customers[k] = v
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(customers)
	}
}

// deleteCustomer deletes a customer
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, ok := customers[id]; ok {
		delete(customers, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(customers)
	}
}

func main() {

	//ca := createCustomer("1", "Andy", "Developer", "S", "S", true)
	//cb := createCustomer("2", "Peter", "Developer", "S", "S", true)
	//customers[ca.ID] = *ca
	//customers[cb.ID] = *cb

	fileServer := http.FileServer(http.Dir("./static"))

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	router.Handle("/", fileServer)

	fmt.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}