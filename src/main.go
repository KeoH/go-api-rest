package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func IndexView(w http.ResponseWriter, r *http.Request){
	fmt.Println("Hooola en la consola")
	fmt.Fprintf(w, "Hoooola")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", IndexView).Methods("GET")

	router.HandleFunc("/tasks", AllTaskEndPoint).Methods("GET")
	router.HandleFunc("/tasks", CreateTaskEndpoint).Methods("POST")
	router.HandleFunc("/tasks/{id}", GetTaskEndpoint).Methods("GET")
	router.HandleFunc("/tasks/{id}", UpdateTaskEndpoint).Methods("PUT")
	router.HandleFunc("/tasks/{id}", DeleteTaskEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}