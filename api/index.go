package api

import (
	"net/http"
)

// Handler is exported for use as a vercel serverless function
func Handler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello"))
}
