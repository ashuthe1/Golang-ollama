package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"FealtyX/handlers"
	"FealtyX/ollama"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
	r.HandleFunc("/students", handlers.GetStudents).Methods("GET")
	r.HandleFunc("/students/{id}", handlers.GetStudent).Methods("GET")
	r.HandleFunc("/students/{id}", handlers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", handlers.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/students/{id}/summary", ollama.GetStudentSummary).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
