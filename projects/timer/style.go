package main

import "github.com/charmbracelet/lipgloss"

/***
 * Lipgloss
 ***/

type styles struct {
	app       lipgloss.Style
	appWidth  int
	appHeight int
	core      lipgloss.Style
}

func initStyles() *styles {
	return &styles{
		app: lipgloss.NewStyle().
			Padding(1, 2).
			Margin(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("13")),

		appWidth:  80,
		appHeight: 24,

		core: lipgloss.NewStyle().
			Align(lipgloss.Center),
	}
}

func (s *styles) Resize(x, y int) {
	width := x - s.app.GetHorizontalMargins() - s.app.GetHorizontalBorderSize()
	height := y - s.app.GetVerticalMargins() - s.app.GetVerticalBorderSize()
	s.app = s.app.Width(width).Height(height)
	s.appWidth = x - s.app.GetHorizontalFrameSize()
	s.appHeight = y - s.app.GetVerticalFrameSize()
	s.core = s.core.Width(s.appWidth)
}
