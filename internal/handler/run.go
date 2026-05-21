package handler

import (
	"encoding/json"
	"net/http"
)

type RunRequest struct {
	Language string `json:"language"`
	Source   string `json:"source"`
	Tests    []Test `json:"tests"`
}

type Test struct {
	Stdin          string `json:"stdin"`
	ExpectedStdout string `json:"expected_stdout"`
}

type TestResult struct {
	Status string `json:"status"`
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

type RunResponse struct {
	Status string       `json:"status"`
	Tests  []TestResult `json:"tests"`
}

func RunHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if req.Language == "" || req.Source == "" {
		http.Error(w, "language and source required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RunResponse{
		Status: "accepted",
		Tests:  []TestResult{{Status: "accepted", Stdout: "ok"}},
	})
}
