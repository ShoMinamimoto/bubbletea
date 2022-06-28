package main

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	innerWidth  int
	innerHeight int

	app lipgloss.Style
}

func newStyles() *Styles {
	//simple border around the app
	app := lipgloss.NewStyle().
		BorderStyle(lipgloss.ThickBorder()).
		Padding(1, 2).
		Margin(1, 2)

	//set reasonable default dimensions
	innerWidth := 80 - app.GetHorizontalFrameSize()
	innerHeight := 24 - app.GetVerticalFrameSize()

	return &Styles{
		innerWidth:  innerWidth,
		innerHeight: innerHeight,
		app:         app,
	}
}

func (s *Styles) Resize(x, y int) {
	s.innerWidth = x - s.app.GetHorizontalFrameSize()
	s.innerHeight = y - s.app.GetVerticalFrameSize()
}

type TaskStyles struct {
	NormalLabel   lipgloss.Style
	NormalTimer   lipgloss.Style
	SelectedLabel lipgloss.Style
	SelectedTimer lipgloss.Style
	ActiveLabel   lipgloss.Style
	ActiveTimer   lipgloss.Style
}

func NewTaskStyles() (s TaskStyles) {
	s.NormalLabel = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
		Padding(0, 0, 0, 2)

	s.NormalTimer = s.NormalLabel.Copy().Padding(0, 2, 0, 0)

	s.SelectedLabel = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#f792ff", Dark: "#ad58b4"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#ee6ff8", Dark: "#ee6ff8"}).
		Padding(0, 0, 0, 1)

	s.SelectedTimer = s.SelectedLabel.Copy().Padding(0, 1, 0, 0).
		Border(lipgloss.NormalBorder(), false, true, false, false)

	s.ActiveLabel = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#6F73F8", Dark: "#6F73F8"}).
		Padding(0, 0, 0, 2)

	s.ActiveTimer = s.ActiveLabel.Copy().Padding(0, 2, 0, 0)

	return s
}
