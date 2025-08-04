package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", AddHandler())
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}