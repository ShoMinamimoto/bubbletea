package main

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	quit      key.Binding
	forceQuit key.Binding
	help      key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.quit,
		k.help,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.quit,
			k.help,
		},
	}
}

func newKeyMap() *KeyMap {
	return &KeyMap{
		quit: key.NewBinding(
			key.WithKeys("q"),
			key.WithHelp("q", "quit"),
		),
		forceQuit: key.NewBinding(
			key.WithKeys("ctrl+c"),
		),
		help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
	}
}

type TaskKeyMap struct {
	start key.Binding
	stop  key.Binding
	edit  key.Binding
}

func NewTaskKeyMap() *TaskKeyMap {
	return &TaskKeyMap{
		start: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "start timer"),
		),
		stop: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "stop timer"),
			key.WithDisabled(),
		),
		edit: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "edit task"),
		),
	}
}
