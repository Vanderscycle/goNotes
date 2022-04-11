package indexPage

import (
	"goNotes/keymaps"

	"github.com/charmbracelet/bubbles/key"
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
	keymap     keymaps.KeyMap
	title      string
	paragraph  string
	keyPressed string
}

func PageInitialModel() Page {
	return Page{
		keymap:     keymaps.DefaultKeyMap,
		title:      "Go Notes",
		paragraph:  "Go TUI build using skate.sh tools to wrap around Task Warrior excellent CLI",
		keyPressed: ""}
}

func (m Page) Init() tea.Cmd {
	// Initialize sub-models
	return nil
}

func (m Page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keymap.Down):
			m.keyPressed = "down"
			// fmt.Print(m.keyPressed)
			return m, nil

		case key.Matches(msg, m.keymap.Up):
			m.keyPressed = "up"
			// log.Print(m.keyPressed)
			return m, nil
		}
	}
	return m, nil

}

func (m Page) View() string {
	s := lipgloss.JoinVertical(lipgloss.Center, title.Render(m.title), paragraph.Render(m.paragraph))
	s += lipgloss.NewStyle().Render(m.keyPressed)
	return s
}
