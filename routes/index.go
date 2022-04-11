package indexPage

//figure out how to use sub directories as I want to have more than one route
//https://stackoverflow.com/questions/24763347/golang-subdirectories
// figureout how to use github imports
import (
	"goNotes/keymaps"
	taskwarrior "goNotes/taskWarrior"
	"log"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	title     = lipgloss.NewStyle().Align(lipgloss.Center).Padding(2)
	paragraph = lipgloss.NewStyle().
			Align(lipgloss.Center).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63")).Padding(2)
)

type Page struct {
	keymap     keymaps.KeyMap
	title      string
	paragraph  string
	keyPressed string
	// command    taskwarrior.Command
}

func PageInitialModel() Page {
	defaultList := taskwarrior.Cmds[1]
	// defaultListDisplay, err := taskwarrior.Api(defaultList)
	// if err != nil {
	// 	panic(err)
	// }
	log.Print(defaultList)
	return Page{
		keymap:     keymaps.DefaultKeyMap,
		title:      "Go Notes",
		paragraph:  "Go TUI build using skate.sh tools to wrap around Task Warrior excellent CLI",
		keyPressed: "",
		// command:    defaultList}
	}
}

func (m Page) Init() tea.Cmd {
	// Initialize sub-models
	return nil
}

func (m Page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keymap.Down):
			m.keyPressed = "down"
			// fmt.Print(m.keyPressed)
			return m, nil

		case key.Matches(msg, m.keymap.Up):
			m.keyPressed = "up"
			// log.Print(m.keyPressed)
			return m, nil
		}
	}
	return m, nil

}

func (m Page) View() string {
	s := lipgloss.JoinVertical(lipgloss.Center, title.Render(m.title), paragraph.Render(m.paragraph))
	s += lipgloss.NewStyle().Render(m.keyPressed)
	s += lipgloss.NewStyle().Render(taskwarrior.Cmds[1].FilterValue())
	return s
}
