package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func handleRequest() {
	http.HandleFunc("/", main_page_handler)
	http.HandleFunc("/create_aim", create_aim)
	http.HandleFunc("/save_article", save_article)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()

}
