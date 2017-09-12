package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// runCommand is running a command and printing the stderr and stdout to stdout.
func runCommand(command string, args []string, ignoreErrors, verbose bool) {
	cmd := exec.Command(command, args...)
	if verbose {
		fmt.Printf("Running command: %s %s\n", command, strings.Join(args, " "))
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open stdout pipe", err)
		os.Exit(1)
	}

	stdoutScanner := bufio.NewScanner(stdout)
	go func() {
		for stdoutScanner.Scan() {
			fmt.Println(stdoutScanner.Text())
		}
	}()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open stderr pipe", err)
		os.Exit(1)
	}

	stderrScanner := bufio.NewScanner(stderr)
	go func() {
		for stderrScanner.Scan() {
			fmt.Println(stderrScanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error running command", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s exited with errors: %s\n", command, err.Error())

		if !ignoreErrors {
			os.Exit(1)
		}
	}
}
