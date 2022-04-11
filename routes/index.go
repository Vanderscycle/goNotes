package indexPage

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	title     = lipgloss.NewStyle().Align(lipgloss.Center).Padding(2)
	paragraph = lipgloss.NewStyle().
			Align(lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63")).Padding(2)
)

type Page struct {
	title     string
	paragraph string
}

func PageInitialModel() Page {
	return Page{
		title:     "Go Notes",
		paragraph: "Go TUI build using skate.sh tools to wrap around Task Warrior excellent CLI"}
}

func (m Page) Init() tea.Cmd {
	// Initialize sub-models
	return nil
}

func (m Page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Page) View() string {
	s := lipgloss.JoinVertical(lipgloss.Center, title.Render(m.title), paragraph.Render(m.paragraph))
	return s
}
