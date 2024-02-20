package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", initRequest)
	http.ListenAndServe(":8080", nil)
}
