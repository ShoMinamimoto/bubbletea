package main

import "github.com/charmbracelet/lipgloss"

/***
 * Lipgloss
 ***/

type styles struct {
	app         lipgloss.Style
	innerWidth  int
	innerHeight int
	core        lipgloss.Style
}

func initStyles() *styles {
	return &styles{
		app: lipgloss.NewStyle().
			Padding(1, 2).
			Margin(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("13")),

		innerWidth:  80,
		innerHeight: 24,

		core: lipgloss.NewStyle().
			Align(lipgloss.Center),
	}
}

func (s *styles) Resize(x, y int) {
	s.innerWidth = x - s.app.GetHorizontalFrameSize()
	s.innerHeight = y - s.app.GetVerticalFrameSize()
	s.core = s.core.Width(s.innerWidth)
}
