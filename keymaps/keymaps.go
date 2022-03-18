package keymaps

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up   key.Binding
	Down key.Binding
	Quit key.Binding
}

var DefaultKeyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),        // actual keybindings
		key.WithHelp("↑/k", "move up"), // corresponding help text
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "move down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctr+c"),
		key.WithHelp("esq", "quit"),
	),
}