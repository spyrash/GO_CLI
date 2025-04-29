package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func StartShell() {
	reader := bufio.NewReader(os.Stdin)
	for {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		hn, err := os.Hostname()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Print(wd, hn, ">> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate the command and the arguments.
	args := strings.Split(input, " ")

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		// 'cd' to home dir with empty path not yet supported.
		if len(args) < 2 {
			return errors.New("path required")
		}
		// Change the directory and return the error.
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}
	// Pass the program and the arguments separately.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	// Execute the command and return the error.
	return cmd.Run()
	/** TODO: Modify the input indicator:
	add the working directory >> done //ugly
	add the machine’s hostname >> done // ok
	add the current user //
	Browse your input history with the up/down keys*/
}
