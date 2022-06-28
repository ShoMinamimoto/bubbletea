package main

import "github.com/charmbracelet/bubbles/textinput"

type Editor struct {
	labelField textinput.Model
	descField  textinput.Model
	task       *Task
}

func (e Editor) Load(task Task) Editor {
	e.task = &task
	e.labelField.SetValue(e.task.label)
	e.descField.SetValue(e.task.description)

	return e
}

func (e Editor) Save() {
	e.task.label = e.labelField.Value()
	e.task.description = e.descField.Value()
}
