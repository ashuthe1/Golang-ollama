package ollama

import (
	"bytes"
	"encoding/json"
	"net/http"
	"FealtyX/models"
)

func GetStudentSummary(w http.ResponseWriter, r *http.Request, student models.Student) {
	summary, err := generateSummary(student)
	if err != nil {
		http.Error(w, "Failed to generate summary", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"summary": summary})
}

func generateSummary(student models.Student) (string, error) {
	url := "https://api.ollama.com/summarize"
	payload := map[string]interface{}{
		"name":  student.Name,
		"age":   student.Age,
		"email": student.Email,
	}

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer your_ollama_api_key")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result["summary"], nil
}
