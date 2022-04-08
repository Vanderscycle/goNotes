package taskwarrior

import (
	"os"
	"os/exec"
	"syscall"
)

func Api() {

}

func lsExample() {
	const taskArg string = "task"
	userInput := "captuire the user input in bubbleTea"
	binary, lookErr := exec.LookPath(taskArg)
	if lookErr != nil {
		panic(lookErr)
	}

	env := os.Environ()

	args := []string{taskArg, "add", userInput}
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

}
