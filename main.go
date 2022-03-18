package main

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"fmt"
	// "os"
	// "strings"
	// "github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
)

type model struct {
	err      error
	spinner  spinner.Model
	quitting bool
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	return model{spinner: s}
}

func (m *model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
func main() {
	fmt.Printf("hello %s", "test")
	// tea.NewProgram(model tea.Model, opts ...tea.ProgramOption)
}
