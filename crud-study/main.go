package main

import (
	"crud-study/crud-study/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/user/", handlers.UserHandler)

	http.ListenAndServe(":12344", nil)
}
