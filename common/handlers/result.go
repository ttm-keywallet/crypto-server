package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

// Response interface
// swagger:response appJsonResult
type appJsonResult struct {
	// Can be a address
	Result interface{} `json:"result"`
}

// ReturnResult
func ReturnResult(w http.ResponseWriter, r interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(appJsonResult{r})
}

// ReturnResultWithCode
func ReturnResultWithCode(ctx context.Context, w http.ResponseWriter, s int, r interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(s)
	json.NewEncoder(w).Encode(appJsonResult{r})
}
