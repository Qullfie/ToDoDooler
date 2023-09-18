package main

import (
	"net/http"
	"text/template"
	"time"
)

func main_page_handler(w http.ResponseWriter, r *http.Request) {
	task_list := []Task{}
	task_list = append(task_list, Task{"Make site", time.Now(), true})
	task_list = append(task_list, Task{"1 Laba", time.Now(), false})
	task_list = append(task_list, Task{"2 Laba", time.Now(), false})

	tmpl, err := template.ParseFiles("templates/main_page.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, task_list)
}
