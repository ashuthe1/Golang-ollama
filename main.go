package main

import (
	"net/http"
	"strconv"

	"FealtyX/handlers"
	"FealtyX/ollama"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
	r.HandleFunc("/students", handlers.GetStudents).Methods("GET")
	r.HandleFunc("/students/{id}", handlers.GetStudent).Methods("GET")
	r.HandleFunc("/students/{id}", handlers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", handlers.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/students/{id}/summary", getStudentSummary).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func getStudentSummary(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	handlers.StudentsMu.Lock()
	student, exists := handlers.Students[id]
	handlers.StudentsMu.Unlock()

	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	ollama.GetStudentSummary(w, r, student)
}
