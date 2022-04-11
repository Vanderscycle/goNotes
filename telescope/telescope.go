package telescope

//named after the amazing telescope in neovim
import (
	"goNotes/keymaps"
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
	docStyle = lipgloss.NewStyle().Padding(1)
)

type Page struct {
	keymap    keymaps.KeyMap
	title     string
	paragraph string
	list      list.Model
}

func PageInitialModel() Page {
	return Page{
		title:     "commands",
		paragraph: "Go",
		list:      list.New(taskwarrior.Cmds, list.NewDefaultDelegate(), 0, 0),
		keymap:    keymaps.DefaultKeyMap,
	}
}

func (m Page) Init() tea.Cmd {
	// Initialize sub-models
	return nil
}

func (m Page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd tea.Cmd
	)
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch {

		// case key.Matches(msg, m.keymap.Quit):
		// 	return m, tea.Quit

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

		case key.Matches(msg, m.keymap.Status):
			m.list.SetShowStatusBar(!m.list.ShowStatusBar())
			return m, nil

		case key.Matches(msg, m.keymap.Help):
			m.list.SetShowHelp(!m.list.ShowHelp())
			return m, nil
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		m.list, cmd = m.list.Update(msg)
		return m, cmd

	}
	return m, nil
}

func (m Page) View() string {
	s := docStyle.Render(m.list.View())
	return s
}
