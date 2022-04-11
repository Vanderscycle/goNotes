package main

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"fmt"
	"os"
	"time"

	"goNotes/keymaps"
	taskwarrior "goNotes/taskWarrior"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
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
	spinner  spinner.Model
	list     list.Model
	quitting bool
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	m.list.Title = "My Fave Things"
	return tea.Batch(m.spinner.Tick)
}

//update
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keymap.Quit):
			msg2 := "quit"
			fmt.Println("%s", msg2)
			m.quitting = true
			// os.Exit(1)
			return m, tea.Quit

		case key.Matches(msg, m.keymap.Down):
			if m.state == "cmd" {
				fmt.Printf("Down")
				m.list.CursorDown()
			}
			return m, nil

		case key.Matches(msg, m.keymap.Up):
			if m.state == "cmd" {
				fmt.Printf("up")
				m.list.CursorUp()
			}
			return m, nil

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
			oldState := m.state
			switch m.state {
			case "home":
				m.state = "cmd"
			case "loading":
				m.state = "cmd" //not specific
			case "cmd":
				m.state = "home"
			}
			fmt.Printf("State change; previous %s, new: %s", oldState, m.state)
			return m, nil
		default:
			// fmt.Printf("%s && %v", msg, key.Matches(msg, m.keymap.Quit))
			return m, nil
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

//view
func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s Loading forever...press q to quit\n\n", m.spinner.View())
	str2 := "work on the views Henri"
	switch m.state {
	case "home":
		return str2
	case "loading":
		return str
	case "cmd":
		return docStyle.Render(m.list.View())
	}
	if m.quitting {
		return str + "\n"
	}
	//TODO: key detection and state management e.g. if user press ? then we show a list of all cmd, if he press
	//  else {
	// 	return docStyle.Render(m.list.View())
	// }

	return "waiting"
}

func initialModel() model {
	//spinner
	s := spinner.New()
	s.Spinner = spinner.Dot

	return model{spinner: s, keymap: keymaps.DefaultKeyMap, list: list.New(taskwarrior.Cmds, list.NewDefaultDelegate(), 0, 0), state: "home"}
}

func main() {

	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
