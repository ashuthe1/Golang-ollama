package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"FealtyX/models"
)

var (
	Students   = make(map[int]models.Student) // Exported variable
	NextID     = 1
	StudentsMu sync.Mutex // Exported variable
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	StudentsMu.Lock()
	defer StudentsMu.Unlock()

	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	student.ID = NextID
	NextID++
	Students[student.ID] = student

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	StudentsMu.Lock()
	defer StudentsMu.Unlock()

	var result []models.Student
	for _, student := range Students {
		result = append(result, student)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	StudentsMu.Lock()
	defer StudentsMu.Unlock()

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	student, exists := Students[id]
	if !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	StudentsMu.Lock()
	defer StudentsMu.Unlock()

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedStudent models.Student
	if err := json.NewDecoder(r.Body).Decode(&updatedStudent); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, exists := Students[id]; !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	updatedStudent.ID = id
	Students[id] = updatedStudent

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedStudent)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	StudentsMu.Lock()
	defer StudentsMu.Unlock()

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, exists := Students[id]; !exists {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	delete(Students, id)
	w.WriteHeader(http.StatusNoContent)
}
