package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"time"
)

/***
 * Key + Help bubbles
 ***/

type keyMap struct {
	exampleKey     key.Binding
	stopwatchStart key.Binding
	stopwatchStop  key.Binding
	stopwatchReset key.Binding
	help           key.Binding
	quit           key.Binding
}

func initKeyMap() *keyMap {
	return &keyMap{
		exampleKey: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "do a thing"),
		),
		stopwatchStart: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "start watch"),
		),
		stopwatchStop: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "stop watch"),
			key.WithDisabled(),
		),
		stopwatchReset: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "reset watch"),
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
		k.stopwatchStart,
		k.stopwatchStop,
		k.stopwatchReset,
		k.help,
		k.quit,
	}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.exampleKey},
		{k.stopwatchStart, k.stopwatchStop, k.stopwatchReset},
		{k.help, k.quit},
	}
}

/***
 * Model definitions
 ***/

// model for the entire program
type model struct {
	// insert global variables here
	stopwatch   stopwatch.Model
	styles      *styles
	keys        *keyMap
	help        help.Model
	testMessage string
}

// returns a model with default values
func initialModel() model {
	return model{
		stopwatch:   stopwatch.NewWithInterval(time.Second),
		styles:      initStyles(),
		keys:        initKeyMap(),
		help:        help.New(),
		testMessage: "Hello World!",
	}
}

// Init returns a starting command or nil
func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
	)
}

// Update consumes messages and returns an updated model and command
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// this narrows down msg's type
	switch msg := msg.(type) {

	// respond to resizing
	case tea.WindowSizeMsg:
		m.styles.Resize(msg.Width, msg.Height)
		m.help.Width = m.styles.appWidth

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.exampleKey):
			m.testMessage = "I did a thing!"
		case key.Matches(msg, m.keys.stopwatchStart, m.keys.stopwatchStop):
			m.keys.stopwatchStart.SetEnabled(m.stopwatch.Running())
			m.keys.stopwatchStop.SetEnabled(!m.stopwatch.Running())
			return m, m.stopwatch.Toggle()
		case key.Matches(msg, m.keys.stopwatchReset):
			return m, m.stopwatch.Reset()
		case key.Matches(msg, m.keys.help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.quit):
			return m, tea.Quit
		}

	case stopwatch.TickMsg, stopwatch.StartStopMsg, stopwatch.ResetMsg:
		var cmd tea.Cmd
		m.stopwatch, cmd = m.stopwatch.Update(msg)
		return m, cmd
	}
	return m, nil
}

// View returns a string that contains the entire display
func (m model) View() string {
	// render test message
	testMsg := m.styles.core.Render(m.testMessage) + "\n"
	stopwatchView := m.styles.core.Render(m.stopwatch.View()) + "\n"

	// generate Help view
	helpView := lipgloss.Place(m.styles.appWidth,
		m.styles.appHeight-lipgloss.Height(testMsg+stopwatchView)+1,
		lipgloss.Right,
		lipgloss.Bottom,
		m.help.View(m.keys))

	return m.styles.app.Render(testMsg + stopwatchView + helpView)
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
