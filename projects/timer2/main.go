package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	program := tea.NewProgram(initModel(), tea.WithAltScreen())
	if err := program.Start(); err != nil {
		fmt.Printf("Whoops: %v", err)
		os.Exit(1)
	}
}

type keyMap struct {
	quit    key.Binding
	help    key.Binding
	newTask key.Binding
}

func initKeyMap() keyMap {
	return keyMap{
		quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		newTask: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add new task"),
		),
	}
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.quit,
		k.help,
	}
}

func (k keyMap) FullHelp() [][]key.Binding {
	//TODO implement me
	panic("implement me")
}

type Model struct {
	state    uint
	tasks    []Task
	taskList list.Model
	keys     keyMap
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	}

	//pass to list Model
	newTaskList, taskListCmd := m.taskList.Update(msg)
	m.taskList = newTaskList
	cmds = append(cmds, taskListCmd)

	//default
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	//TODO implement me
	panic("implement me")
}

func initModel() Model {
	var tasks []Task
	return Model{
		taskList: list.New(tasks, initTaskDelegate(), 40, 20),
	}
}
