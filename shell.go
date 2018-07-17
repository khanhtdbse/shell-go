package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("KhanhShell> ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(os.Stderr, err)
		}

		fullCmd := strings.TrimSuffix(input, "\n")
		args := strings.Split(fullCmd, " ")

		execInput(args[0], args[1:])
	}
}

func execInput(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
