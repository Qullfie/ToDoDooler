package main

import (
	"net/http"
)

func handleRequest() {
	http.HandleFunc("/", main_page_handler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
