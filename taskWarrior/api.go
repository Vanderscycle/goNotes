package taskwarrior

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/charmbracelet/bubbles/list"
)

type Command struct {
	title string
	desc  string
}

func (c Command) Title() string       { return c.title }
func (c Command) Description() string { return c.desc }
func (c Command) FilterValue() string { return c.title }

const TaskArg string = "task"

// := can only be used inside functions
var Cmds = []list.Item{
	Command{title: "add", desc: "add a task"},
	Command{title: "list", desc: "list all tasks"},
	Command{title: "delete", desc: "delete's a task"}, //need to pass a value
}

func Api(c *Command) {
	userInput := "capture the user input in bubbleTea"
	binary, lookErr := exec.LookPath(TaskArg)
	if lookErr != nil {
		log.Fatal("TaskWarrior not installed")
		panic(lookErr)
	}

	env := os.Environ()

	args := []string{TaskArg, c.Title(), userInput}
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
