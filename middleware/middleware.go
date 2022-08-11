package middleware

import (
	"encoding/json"
	"net/http"
	"os"
)

func requestIDHandler(next http.Handler) http.Handler {
	api_key := os.Getenv("X_API_KEY")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("x-api-key")
		if api_key != requestID {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			resp := make(map[string]string)
			resp["Message"] = "Unauthorized access without api key"
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
			return
		}
		w.Header().Set("x-api-key", requestID)
		next.ServeHTTP(w, r)
	})
}