package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	appnameStyle = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0,1)
	faintStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)
	numberStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
	menuStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000"))
	menuStyle2 = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff00"))
)

func (m model) View() string {

	s := appnameStyle.Render("NOTES APP") + "\n\n"

	if m.state == titleView {
		s += "Note Title: \n\n"
		s += m.textinput.View() + "\n\n"
		s += menuStyle2.Render("enter - continue, esc - discard")
	}

	if m.state == bodyView {
		s += "Note: \n\n"
		s += m.textarea.View() + "\n\n"
		s += menuStyle2.Render("ctrl+s - save, esc - discard")
	}
	
	if m.state == listView {
		for i, n := range m.notes {
			prefix := " "
			if i == m.listIndex {
				prefix = ">>"
			}

			shortBody := strings.ReplaceAll(n.Body, "\n", " ")

			if len(shortBody) > 30 {
				shortBody = shortBody[:20]
			}

			s += numberStyle.Render(prefix) + n.Title + " | " + faintStyle.Render(shortBody) + "\n\n"
		}

		s += menuStyle.Render("n - new note, q - quit")
	}

	return s
}