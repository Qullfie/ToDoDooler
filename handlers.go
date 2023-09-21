package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"text/template"
)

func main_page_handler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/main_page.html")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/tododooler")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//выборка данных

	tasks := []Task{}
	sql_query := fmt.Sprintf("SELECT * from task")
	res, err := db.Query(sql_query)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var task Task
		err = res.Scan(&task.Id, &task.Aim, &task.Solved)
		tasks = append(tasks, task)
	}
	tmpl.Execute(w, tasks)
}

func create_aim(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/create_aim.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}

func save_aim(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/tododooler")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	aim := r.FormValue("aim")
	sql_query := fmt.Sprintf("INSERT INTO task (aim, solved) VALUES ('%s', %d)", aim, 0)
	db.Query(sql_query)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func solve_task(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/tododooler")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	task_id := r.FormValue("taskID")
	fmt.Println(task_id)
	sql_query := fmt.Sprintf("UPDATE task SET solved = 1 WHERE id = %v", task_id)
	db.Query(sql_query)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func delete_task(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/tododooler")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	task_id := r.FormValue("taskID")
	fmt.Println(task_id)
	sql_query := fmt.Sprintf("DELETE FROM task WHERE id = %v", task_id)
	db.Query(sql_query)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
