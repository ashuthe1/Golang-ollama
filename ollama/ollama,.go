package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"FealtyX/handlers"
	"FealtyX/models"

	"github.com/gorilla/mux"
)

const (
	googlePalm2Endpoint = "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent"
	googleApiKey        = "AIzaSyDXFNc6-NpTLqzmHA6PZvLK9ctDEffz7rE" // Replace with your actual Google API key
)

func GetStudentSummary(w http.ResponseWriter, r *http.Request) {
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

	// summary, err := generateSummaryWithllama2(student)
	summary, err := generateSummaryWithGooglePalm2(student)
	if err != nil {
		http.Error(w, "Failed to generate summary", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"summary": summary})
}

func generateSummaryWithllama2(student models.Student) (string, error) {
	msg := fmt.Sprint("Generate Summary for this student profile\n", "Name: ", student.Name, "\nAge: ", student.Age, "\nEmail: ", student.Email)
	fmt.Println(msg)
	url := "http://localhost:11434/api/generate"
	payload := map[string]interface{}{
		"content": msg,
		"model":   "llama3",
	}

	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

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
func generateSummaryWithGooglePalm2(student models.Student) (string, error) {
	msg := fmt.Sprintf("Generate Summary for this student profile Name: %s, Age: %d, Email: %s", student.Name, student.Age, student.Email)
	fmt.Println(msg)

	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{
						"text": msg,
					},
				},
			},
		},
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s?key=%s", googlePalm2Endpoint, googleApiKey)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Google PaLM API returned non-200 status: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	// Extract the summary from the response
	candidates, ok := result["candidates"].([]interface{})
	if !ok || len(candidates) == 0 {
		return "", fmt.Errorf("Failed to extract candidates from response")
	}

	content, ok := candidates[0].(map[string]interface{})["content"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("Failed to extract content from response")
	}

	parts, ok := content["parts"].([]interface{})
	if !ok || len(parts) == 0 {
		return "", fmt.Errorf("Failed to extract parts from content")
	}

	summary, ok := parts[0].(map[string]interface{})["text"].(string)
	if !ok {
		return "", fmt.Errorf("Failed to extract summary text from parts")
	}

	return summary, nil
}
