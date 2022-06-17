package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

func main() {
	p := tea.NewProgram(newModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Something went wrong: %v", err)
		os.Exit(1)
	}
}

type Model struct {
	state uint8
	keys  *KeyMap
	style *Styles
	help  help.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.style.Resize(msg.Width, msg.Height)
		m.help.Width = m.style.innerWidth

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var (
		sections        []string
		availableHeight = m.style.innerHeight
	)

	helpSection := lipgloss.Place(
		m.style.innerWidth,
		availableHeight,
		lipgloss.Left,
		lipgloss.Bottom,
		m.help.View(m.keys),
	)
	sections = append(sections, helpSection)

	content := lipgloss.JoinVertical(lipgloss.Left, sections...)

	return m.style.app.Render(content)
}

func newModel() Model {
	return Model{
		state: 0,
		keys:  newKeyMap(),
		style: newStyles(),
		help:  help.New(),
	}
}
