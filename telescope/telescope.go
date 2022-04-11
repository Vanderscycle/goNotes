package telescope

//named after the amazing telescope in neovim
import (
	"goNotes/keymaps"
	taskwarrior "goNotes/taskWarrior"
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	title = lipgloss.NewStyle().
		Align(lipgloss.Center).
		Padding(2).
		Foreground(lipgloss.Color("170"))
	paragraph = lipgloss.NewStyle().
			Align(lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63")).Padding(2)
	docStyle = lipgloss.NewStyle().Padding(1)
)

type Page struct {
	keymap    keymaps.KeyMap
	paragraph string
	list      list.Model
	searching bool
}

func PageInitialModel() Page {
	l := list.New(taskwarrior.Cmds, list.NewDefaultDelegate(), 0, 0)
	l.Styles.Title = title
	l.Title = "commands"

	return Page{
		paragraph: "Go",
		list:      l,
		keymap:    keymaps.DefaultKeyMap,
		searching: false}
}

func (m Page) Init() tea.Cmd {
	// Initialize sub-models
	//TODO: css
	// l.Styles.PaginationStyle = paginationStyle
	// l.Styles.HelpStyle = helpStyle
	return tea.EnterAltScreen
}

func (m Page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		// m.list, cmd = m.list.Update(msg)
		return m, cmd
	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			log.Print("filter")
			m.searching = true
			break
		} else {
			m.searching = false
		}

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

	}
	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Page) View() string {
	s := docStyle.Render(m.list.View())
	return s
}
