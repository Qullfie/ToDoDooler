package main

import "time"

type Task struct {
	Name     string
	DateTime time.Time
	Solved   bool
}

func (t *Task) ChangeName(new_name string) {
	t.Name = new_name
}

func (t *Task) SolveTask() {
	t.Solved = true
}
