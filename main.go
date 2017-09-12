package main

import (
	"flag"
	"fmt"
	"os"
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

	args = append(args, *packageName)

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

	runCommand("gometalinter", args, *ignoreErrors, *verbose)
}
