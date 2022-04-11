package main

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

// we will need to use channels (to check the state of the app)
import (
	"fmt"
	"os"
	"time"

	"goNotes/keymaps"
	indexPage "goNotes/routes"
	"goNotes/telescope"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// const listHeight Int = 14
var state = []string{"cmd", "home"}

type errMsg error

type tickMsg time.Time

//tea model
type model struct {
	keymap   keymaps.KeyMap
	index    tea.Model
	err      error
	list     tea.Model
	state    string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.index.Init(), m.list.Init())
}

//update
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch {

		case key.Matches(msg, m.keymap.Quit):
			msg2 := "quit"
			fmt.Println("%s", msg2)
			m.quitting = true
			// os.Exit(1)
			cmds = append(cmds, tea.Quit)

		case key.Matches(msg, m.keymap.State):
			switch m.state {
			case "home":
				m.state = "cmd"
			case "loading":
				m.state = "cmd" //not specific
			case "cmd":
				m.state = "home"
			}
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.index, cmd = m.index.Update(msg)
	cmds = append(cmds, cmd)
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

//view
func (m model) View() string {
	if m.err != nil {
		//TODO: catch the error , display and return to the home screen
		return m.err.Error()
	}
	str := fmt.Sprintf("[MSG]: %s", "Quitting but turn this into a func? with log warning/or get a logger")
	if m.quitting {
		return str + "\n"
	}

	switch m.state {
	case "home":
		return m.index.View()
	case "cmd":
		return m.list.View()
	}
	return "waiting"
}

func initialModel() model {

	return model{
		index:    indexPage.PageInitialModel(),
		keymap:   keymaps.DefaultKeyMap,
		err:      nil,
		state:    "cmd",
		list:     telescope.PageInitialModel(),
		quitting: false,
	}
}

func main() {

	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
