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
