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

var myCuteBorder = lipgloss.Border{
	Top:         "._.:*:",
	Bottom:      "._.:*:",
	Left:        "|*",
	Right:       "|*",
	TopLeft:     "*",
	TopRight:    "*",
	BottomLeft:  "*",
	BottomRight: "*",
}

type styles struct {
	app       lipgloss.Style
	appTop    lipgloss.Style
	appBottom lipgloss.Style
	appwidth  int
	appheight int
	core      lipgloss.Style
}

func initStyles() *styles {
	border := lipgloss.HiddenBorder()
	return &styles{
		app: lipgloss.NewStyle().
			Padding(0, 2).
			Margin(0, 2).
			Border(border, false, true),

		appTop: lipgloss.NewStyle().
			Padding(1, 2, 0).
			Margin(1, 2, 0).
			Border(border, true, true, false),

		appBottom: lipgloss.NewStyle().
			Padding(0, 2, 1).
			Margin(0, 2, 1).
			Border(border, false, true, true),

		appwidth:  80,
		appheight: 24,

		core: lipgloss.NewStyle().
			Width(50).
			Align(lipgloss.Center),
	}
}

func (s *styles) Resize(x, y int) {
	s.appwidth = x - s.app.GetHorizontalFrameSize()
	s.appheight = y - s.appTop.GetVerticalFrameSize() - s.appBottom.GetVerticalFrameSize()
	s.app = s.app.
		Width(s.appwidth + s.app.GetHorizontalPadding())
	s.core = s.core.Width(s.appwidth)
}

func (s styles) RenderFlag(flag, str string) string {
	lines := strings.SplitAfter(str, "\n")
	colors := flags[flag].colors
	sectionHeight := (s.appheight + s.app.GetVerticalBorderSize()) / len(colors)
	var sections []string
	for index, color := range colors {
		var nextSection string
		switch index {
		case 0:
			next := lines[0 : sectionHeight-1]
			nextSection = s.appTop.
				BorderBackground(lipgloss.Color(color)).
				Render(strings.Join(next, ""))
		case len(colors) - 1:
			next := lines[index*sectionHeight-1:]
			nextSection = s.appBottom.
				BorderBackground(lipgloss.Color(color)).
				Render(strings.Join(next, ""))
		default:
			next := lines[index*sectionHeight-1 : (index+1)*sectionHeight-1]
			nextSection = s.app.
				BorderBackground(lipgloss.Color(color)).
				Render(strings.Join(next, ""))
		}
		sections = append(sections, nextSection)
	}
	return strings.Join(sections, "")
}

type prideFlag struct {
	colors []string
}

var flags = map[string]prideFlag{
	"trans": {
		colors: []string{"#5BCEFA", "#F5A9B8", "#FFFFFF", "#F5A9B8", "#5BCEFA"},
	},
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
	currentFlag string
	keys        *keyMap
	help        help.Model
	testMessage string
}

// returns a model with default values
func initialModel() model {
	return model{
		styles:      initStyles(),
		currentFlag: "trans",
		keys:        newKeyMap(),
		help:        help.New(),
		testMessage: "Happy Pride!",
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
		m.styles.Resize(msg.Width, msg.Height)
		m.help.Width = m.styles.appwidth

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
	availableHeight := m.styles.appheight

	// render test message
	testMsg := m.styles.core.Render(m.testMessage)
	availableHeight -= lipgloss.Height(testMsg)

	// generate Help view
	helpView := lipgloss.Place(m.styles.appwidth, availableHeight, lipgloss.Bottom, lipgloss.Right, m.help.View(m.keys))

	return m.styles.RenderFlag(m.currentFlag, testMsg+"\n"+helpView)
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
