package main

type Task struct {
	Id     uint
	Aim    string
	Solved bool
}

func (t *Task) ChangeAim(new_aim string) {
	t.Aim = new_aim
}

func (t *Task) SolveTask() {
	t.Solved = true
}
