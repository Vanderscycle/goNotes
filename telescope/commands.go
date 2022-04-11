package indexPage

import (
	taskwarrior "goNotes/taskWarrior"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
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
	list      list.Model
	keymap    keymaps.KeyMap
}

func PageInitialModel() Page {
	return Page{
		title:     "commands",
		paragraph: "Go",
		list:      list.New(taskwarrior.Cmds, list.NewDefaultDelegate(), 0, 0),
	}
}

func (m Page) Init() tea.Cmd {
	// Initialize sub-models
	return nil
}

func (m Page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keymap.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keymap.Down):
			m.list.CursorDown()
			return m, nil

		case key.Matches(msg, m.keymap.Up):
			m.list.CursorUp()
			return m, nil

		case key.Matches(msg, m.keymap.Search):
			v := !m.list.ShowTitle()
			m.list.SetShowTitle(v)
			m.list.SetShowFilter(v)
			m.list.SetFilteringEnabled(v)
			return m, nil

		case key.Matches(msg, m.keymap.Help):
			m.list.SetShowHelp(!m.list.ShowHelp())
			return m, nil

		}
	default:
		return m, nil
	}
	return m, nil

}

func (m Page) View() string {
	s := lipgloss.JoinVertical(lipgloss.Center, title.Render(m.title), paragraph.Render(m.paragraph))
	return s
}
