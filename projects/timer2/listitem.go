package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"io"
)

type Task struct {
	label   string
	details string
	timer   stopwatch.Model
}

func (t Task) FilterValue() string {
	return ""
}

type delegateKeyMap struct {
	startTimer key.Binding
	stopTimer  key.Binding
	editTask   key.Binding
}

type taskDelegate struct {
	keys        delegateKeyMap
	timerActive bool
	styles      list.DefaultItemStyles
}

func initTaskDelegate() taskDelegate {
	return taskDelegate{
		keys: delegateKeyMap{
			startTimer: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "start timer"),
			),
			stopTimer: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "stop timer"),
				key.WithDisabled(),
			),
			editTask: key.NewBinding(
				key.WithKeys("e"),
				key.WithHelp("e", "edit task"),
			),
		},
		timerActive: false,
		styles:      list.NewDefaultItemStyles(),
	}
}

func (t taskDelegate) ShortHelp() []key.Binding {
	return []key.Binding{
		t.keys.startTimer,
		t.keys.stopTimer,
		t.keys.editTask,
	}
}

func (t taskDelegate) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		t.ShortHelp(),
	}
}

func (t taskDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	//TODO implement me
	panic("implement me")
}

func (t taskDelegate) Height() int {
	return 1
}

func (t taskDelegate) Spacing() int {
	return 1
}

func (t taskDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, t.keys.startTimer, t.keys.stopTimer):
			return m.SelectedItem().(Task).timer.Toggle()
		case key.Matches(msg, t.keys.editTask):
		}
	}
	return nil
}
