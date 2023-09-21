package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func handleRequest() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", main_page_handler)
	http.HandleFunc("/create_aim", create_aim)
	http.HandleFunc("/save_aim", save_aim)
	http.HandleFunc("/solve_task", solve_task)
	http.HandleFunc("/delete_task", delete_task)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()

}
