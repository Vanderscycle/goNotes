package keymaps

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Quit   key.Binding
	Enter  key.Binding
	State  key.Binding
	Search key.Binding
	Help   key.Binding
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
	Enter:  key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
	State:  key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "change the state of the app")),
	Search: key.NewBinding(key.WithKeys("/"), key.WithHelp("/", "search")),
	Help:   key.NewBinding(key.WithKeys("h"), key.WithHelp("h", "display hides the help menu")),
}
