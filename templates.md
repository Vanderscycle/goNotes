# Templates 

## Simple component page
```go
package indexPage

import tea "github.com/charmbracelet/bubbletea"

type Page struct {
	title     string
	paragraph string
}

func PageInitialModel() Page {
	return Page{
		title:     "Go Notes",
		paragraph: "Go TUI build using skate.sh tools to wrap around Task Warrior excellent CLI"}
}

func (m Page) Init() tea.Cmd {
	// Initialize sub-models
	return nil
}

func (m Page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Page) View() string {
	s := m.title
	s += m.paragraph
	return s
}

```
