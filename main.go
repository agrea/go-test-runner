package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	disableGoTest       = flag.Bool("disable-go-test", false, "Disable go test execution")
	disableGometalinter = flag.Bool("disable-gometalinter", false, "Disable gometalinter checks")
	goTestFlags         = flag.String("go-test-flags", "", "Send custom flags to go test as parameters")
	goTestShort         = flag.Bool("go-test-short", false, "Enable -short mode for go test")
	gometalinterConfig  = flag.String("gometalinter-config", "", "Path to configuration file for gometalinter")
	gometalinterFlags   = flag.String("gometalinter-flags", "", "Send custom flags to gometalinter")
	gometalinterPath    = flag.String("gometalinter-path", ".", "Path for gometalinter to lint. Set it to ./... for recursion")
	ignoreErrors        = flag.Bool("ignore-errors", false, "Continue with the next check on errors")
	packageName         = flag.String("package", "", "Package name to test")
	recursive           = flag.Bool("recursive", false, "Run tests recursivly")
	verbose             = flag.Bool("verbose", false, "Enable verbose output")
)

func main() {
	flag.Parse()

	if *packageName == "" {
		fmt.Println("Missing required parameter -package")
		flag.Usage()
		os.Exit(1)
	}

	if !*disableGometalinter {
		runGometalinter()
	}

	if !*disableGoTest {
		runGoTest()
	}
}

// runGoTest is executing go test with the given parameters.
func runGoTest() {
	fmt.Println("+++ Running go test")

	var args []string

	args = append(args, "test")

	if *goTestFlags != "" {
		args = append(args, *goTestFlags)
	}

	if *goTestShort {
		args = append(args, "-short")
	}

	pkg := *packageName
	if *recursive {
		pkg = fmt.Sprintf("%s/...", pkg)
	}

	args = append(args, pkg)

	runCommand("go", args, *ignoreErrors, *verbose)
}

// runGometalinter is running the gometalinter checks.
func runGometalinter() {
	fmt.Println("+++ Running gometalinter checks")

	var args []string

	if *gometalinterFlags != "" {
		args = append(args, *gometalinterFlags)
	}

	if *gometalinterConfig != "" {
		args = append(args, fmt.Sprintf("--config=%s", *gometalinterConfig))
	}

	args = append(args, *gometalinterPath)

	runCommand("gometalinter.v1", args, *ignoreErrors, *verbose)
}

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
