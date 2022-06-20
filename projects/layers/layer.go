package main

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
)

func layer(top, bottom string) string {
	if lipgloss.Size(top) != lipgloss.Size(bottom) {
		panic("Can't layer differently sized strings")
	}

	b := strings.Builder{}
	topRunes := []rune(top)
	bottomRunes := []rune(bottom)

	for i, v := range topRunes {
		if v == ' ' {
			b.WriteRune(bottomRunes[i])
			continue
		}
		b.WriteRune(v)
	}
	return b.String()
}
