package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"strings"
)

/***
 * Lipgloss
 ***/

var (
	appStyle = lipgloss.NewStyle().
			Padding(1, 2).
			MarginBottom(1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("13"))

	coreStyle = lipgloss.NewStyle().
			Width(50).
			Align(lipgloss.Center)
)

/***
 * Key + Help bubbles
 ***/

type keyMap struct {
	exampleKey key.Binding
	help       key.Binding
	quit       key.Binding
}

func newKeyMap() *keyMap {
	return &keyMap{
		exampleKey: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "do a thing"),
		),
		help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.exampleKey,
		k.help,
	}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.exampleKey},
		{k.help, k.quit},
	}
}

/***
 * Model definitions
 ***/

// model for the entire program
type model struct {
	// insert global variables here
	keys        *keyMap
	help        help.Model
	testMessage string
}

// returns a model with default values
func initialModel() model {
	return model{
		keys:        newKeyMap(),
		help:        help.New(),
		testMessage: "Hello World!",
	}
}

// Init returns a starting command or nil
func (m model) Init() tea.Cmd {
	return nil
}

// Update consumes messages and returns an updated model and command
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// this narrows down msg's type
	switch msg := msg.(type) {

	// respond to resizing
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.exampleKey):
			m.testMessage = "I did a thing!"
		case key.Matches(msg, m.keys.help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

// View returns a string that contains the entire display
func (m model) View() string {
	// render test message
	testMsg := coreStyle.Render(m.testMessage)

	// generate Help view
	helpView := m.help.View(m.keys)

	// calculate whitespace
	padding := 8 - lipgloss.Height(testMsg) - lipgloss.Height(helpView)

	return appStyle.Render(testMsg + strings.Repeat("\n", padding) + helpView)
}

/***
 * main method
 ***/

// main initializes a model and starts a bubbletea program
func main() {
	program := tea.NewProgram(initialModel())
	if err := program.Start(); err != nil {
		fmt.Printf("An error has occurred: %v", err)
		os.Exit(1)
	}
}
