package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

var ErrNoPath = errors.New("path required")

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s >", user.Name)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = runInput(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func runInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return ErrNoPath
		}

		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
