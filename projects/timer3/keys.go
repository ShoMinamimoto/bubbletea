package main

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	quit key.Binding
	help key.Binding
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
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
	}
}
