package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("X_API_KEY")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-api-key")
		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}
		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token")
			return
		}
		c.Next()
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
