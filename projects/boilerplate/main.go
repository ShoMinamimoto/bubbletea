package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

/***
 * Model definitions
 ***/

// model for the entire program
type model struct {
	// insert global variables here
	testMessage string
}

// returns a model with default values
func initialModel() model {
	return model{
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
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl-c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View returns a string that contains the entire display
func (m model) View() string {
	return m.testMessage
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
