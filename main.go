package main

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"fmt"
	"os"
	"time"

	"goNotes/keymaps"
	indexPage "goNotes/routes"
	taskwarrior "goNotes/taskWarrior"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// const listHeight Int = 14
var state = []string{"cmd", "home", "loading"}
var docStyle = lipgloss.NewStyle().Margin(1, 2)

type errMsg error

type tickMsg time.Time

//tea model
type model struct {
	index    tea.Model
	keymap   keymaps.KeyMap
	err      error
	state    string
	list     list.Model
	quitting bool
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.index.Init())
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

		// case key.Matches(msg, m.keymap.Down):
		// 	if m.state == "cmd" {
		// 		m.list.CursorDown()
		// 	}
		// 	return m, nil

		// case key.Matches(msg, m.keymap.Up):
		// 	if m.state == "cmd" {
		// 		m.list.CursorUp()
		// 	}
		// 	return m, nil

		case key.Matches(msg, m.keymap.Search):
			if m.state == "cmd" {
				fmt.Printf("search")
				v := !m.list.ShowTitle()
				m.list.SetShowTitle(v)
				m.list.SetShowFilter(v)
				m.list.SetFilteringEnabled(v)
			}
			return m, nil

		case key.Matches(msg, m.keymap.Help):
			m.list.SetShowHelp(!m.list.ShowHelp())
			return m, nil

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

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		m.list, cmd = m.list.Update(msg)
		return m, cmd

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.index, cmd = m.index.Update(msg)
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
		return docStyle.Render(m.list.View())
	}
	return "waiting"
}

func initialModel() model {

	return model{
		index:    indexPage.PageInitialModel(),
		keymap:   keymaps.DefaultKeyMap,
		err:      nil,
		state:    "home",
		list:     list.New(taskwarrior.Cmds, list.NewDefaultDelegate(), 0, 0),
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
