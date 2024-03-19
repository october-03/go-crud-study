package utils

import "net/http"

func MethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}
