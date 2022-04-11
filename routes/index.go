package index

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	title     string
	paragraph string
}

func initialModel() model {
	return model{
		title:     "Go Notes",
		paragraph: "Go TUI build using skate.sh tools to wrap around Task Warrior excellent CLI"}
}

func (m model) Init() tea.Cmd {
	// Initialize sub-models
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	s := m.title
	s += m.paragraph
	return s
}
