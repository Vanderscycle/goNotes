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
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

type errMsg error

type tickMsg time.Time
type model struct {
	err      error
	progress progress.Model
	spinner  spinner.Model
	quitting bool
}

func initialModel() model {
	//progess bar

	//spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	return model{spinner: s}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n%s\n   %s Loading forever...press q to quit\n\n", "hello", m.spinner.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}

func main() {
	fmt.Printf("hello %s", "test")
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
