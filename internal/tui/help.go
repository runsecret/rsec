package tui

import "github.com/charmbracelet/glamour"

var renderer *glamour.TermRenderer

func RenderMkDown(in string) string {
	r, _ := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
		glamour.WithEmoji(),
	)

	out, _ := r.Render(in)
	return out
}
