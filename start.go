package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonRequest struct {
	Message string `json:"message"`
}

type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req JsonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}

	if req.Message == "" {
		http.Error(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}

	fmt.Println("Received message:", req.Message)

	resp := JsonResponse{
		Status:  "success",
		Message: "Data successfully received",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
