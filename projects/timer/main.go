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

type styles struct {
	app  lipgloss.Style
	core lipgloss.Style
}

func initialStyles() *styles {
	return &styles{
		app: lipgloss.NewStyle().
			Padding(1, 2).
			Margin(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("13")).
			Height(10),

		core: lipgloss.NewStyle().
			Width(80).
			Align(lipgloss.Center),
	}
}

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
	styles      *styles
	keys        *keyMap
	help        help.Model
	testMessage string
}

// returns a model with default values
func initialModel() model {
	return model{
		styles:      initialStyles(),
		keys:        newKeyMap(),
		help:        help.New(),
		testMessage: "Hello World!",
	}
}

// Init returns a starting command or nil
func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

// Update consumes messages and returns an updated model and command
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// this narrows down msg's type
	switch msg := msg.(type) {

	// respond to resizing
	case tea.WindowSizeMsg:
		width := msg.Width - m.styles.app.GetHorizontalMargins() - m.styles.app.GetHorizontalBorderSize()
		height := msg.Height - m.styles.app.GetVerticalMargins() - m.styles.app.GetVerticalBorderSize()
		m.styles.app = m.styles.app.Width(width).Height(height)
		m.help.Width = width

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
	testMsg := m.styles.core.Render(m.testMessage)

	// generate Help view
	helpView := m.help.View(m.keys)
	helpSize := fmt.Sprintf("Help width = %v", lipgloss.Width(helpView))

	// calculate whitespace
	padding := m.styles.app.GetHeight() - m.styles.app.GetVerticalBorderSize() - lipgloss.Height(testMsg) - lipgloss.Height(helpView)

	return m.styles.app.Render(testMsg + strings.Repeat("\n", padding) + helpSize + helpView)
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
