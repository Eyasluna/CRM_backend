package main

import (
	cs "CRM_backend/customer"
	op "CRM_backend/operation"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	ca := cs.CreateCustomer("1", "Andy", "Developer", "S", 320, true)
	cb := cs.CreateCustomer("2", "Peter", "Developer", "S", 408, true)
	cc := cs.CreateCustomer("3", "Andy", "Manager", "S", 47, false)
	op.Customers[ca.ID] = *ca
	op.Customers[cb.ID] = *cb
	op.Customers[cc.ID] = *cc

	fileServer := http.FileServer(http.Dir("./static"))

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/customers", op.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", op.GetCustomer).Methods("GET")
	router.HandleFunc("/customers", op.AddCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", op.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", op.DeleteCustomer).Methods("DELETE")
	router.Handle("/", fileServer)

	fmt.Println("Server is starting on port 3000. You can access it on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
