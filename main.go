package main

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"fmt"
	"os"
	"time"

	// "os"
	// "strings"
	// "github.com/charmbracelet/bubbles/textinput"
	// "github.com/charmbracelet/bubbles/progress"
	// "github.com/charmbracelet/lipgloss"
	"goNotes/keymaps"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// const listHeight Int = 14

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title string
	desc  string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type errMsg error

type tickMsg time.Time

//tea model
type model struct {
	keymap   keymaps.KeyMap
	list     list.Model
	err      error
	spinner  spinner.Model
	quitting bool
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
			os.Exit(1)
			return m, tea.Quit

		case key.Matches(msg, m.keymap.Down):
			fmt.Printf("Down")
			return m, nil

		case key.Matches(msg, m.keymap.Up):
			fmt.Printf("up")
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
	if m.quitting {
		return str + "\n"
	}
	//TODO: key detection and state management e.g. if user press ? then we show a list of all cmd, if he press
	//  else {
	// 	return docStyle.Render(m.list.View())
	// }

	return str
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	m.list.Title = "My Fave Things"
	return m.spinner.Tick
}

func initialModel() model {
	//cmd list
	cmds := []list.Item{
		item{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		item{title: "Nutella", desc: "It's good on toast"},
	}

	//spinner
	s := spinner.New()
	s.Spinner = spinner.Dot

	return model{spinner: s, keymap: keymaps.DefaultKeyMap, list: list.New(cmds, list.NewDefaultDelegate(), 0, 0)}
}

func main() {

	fmt.Printf("hello %s", "test")
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
