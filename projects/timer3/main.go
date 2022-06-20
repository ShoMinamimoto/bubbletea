package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
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
	state    uint8
	keys     *KeyMap
	style    *Styles
	help     help.Model
	taskList list.Model
}

func (m Model) ShortHelp() []key.Binding {
	var keybinds []key.Binding
	keybinds = append(keybinds, m.taskList.ShortHelp()...)
	keybinds = append(keybinds, m.keys.ShortHelp()...)
	return keybinds
}

func (m Model) FullHelp() [][]key.Binding {
	var keybinds [][]key.Binding
	keybinds = append(keybinds, m.taskList.FullHelp()...)
	keybinds = append(keybinds, m.keys.FullHelp()...)
	return keybinds
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.style.Resize(msg.Width, msg.Height)
		m.taskList.SetSize(m.style.innerWidth, m.style.innerHeight-5)
		m.help.Width = m.style.innerWidth

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}

	newTaskList, cmd := m.taskList.Update(msg)
	m.taskList = newTaskList
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var (
		sections        []string
		availableHeight = m.style.innerHeight
	)

	taskSection := m.taskList.View()
	sections = append(sections, taskSection)
	availableHeight -= lipgloss.Height(taskSection)

	helpSection := lipgloss.Place(
		m.style.innerWidth,
		availableHeight,
		lipgloss.Left,
		lipgloss.Bottom,
		m.help.View(m),
	)
	sections = append(sections, helpSection)

	content := lipgloss.JoinVertical(lipgloss.Left, sections...)

	return m.style.app.Render(content)
}

func newModel() Model {
	var testTasks []list.Item
	for i := 1; i <= 5; i++ {
		testTasks = append(testTasks, newTask(i))
	}

	taskList := list.New(testTasks, newTaskDelegate(), 70, 20)
	taskList.SetShowHelp(false)

	return Model{
		state:    0,
		keys:     newKeyMap(),
		style:    newStyles(),
		help:     help.New(),
		taskList: taskList,
	}
}
