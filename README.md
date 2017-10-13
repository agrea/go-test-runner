# Go test runner

The `go-test-runner` is a simple way of running lint checks (via `gometalinter`)
and execute tests (via `go test`) on your projects.

**If you like the project, please star it! If you feel really generous, please
follow [@AgreaIO](https://twitter.com/AgreaIO) on Twitter**

## Installation

### Local installation

    go get -u github.com/agrea/go-test-runner
    go install github.com/agrea/go-test-runner

### Running it via Docker

Below is an example command to run `go-test-runner` via a Docker image. Replace
`<PACKAGE>` with your package path (like `github.com/agrea/go-test-runner`). Any
parameters to `go-test-runner` can be appended to the command.

    docker run -it --rm \
        -v $(pwd):/go/src/<PACKAGE> \
        agrea/go-test-runner:latest \
        go-test-runner \
            -package <PACKAGE>

## Usage and configuration

### Command line instructions

    Usage of ./go-test-runner:
      -disable-go-test
            Disable go test execution
      -disable-gometalinter
            Disable gometalinter checks
      -go-test-flags string
            Send custom flags to go test as parameters
      -go-test-short
            Enable -short mode for go test
      -gometalinter-config string
            Path to configuration file for gometalinter
      -gometalinter-flags string
            Send custom flags to gometalinter
      -gometalinter-path string
            Path for gometalinter to lint. Set it to ./... for recursion (default ".")
      -ignore-errors
            Continue with the next check on errors
      -package string
            Package name to test
      -recursive
            Run tests recursivly
      -verbose
            Enable verbose output

### Gometalinter configuration

We usually run `gometalinter` with a configuration file. You can specify a
Gometalinter configuration file for `go-test-runner` to use with the
`-gometalinter-config` parameter. Here's an example configuration:

    {
      "Enable": [
        "deadcode",
        "vet",
        "gosimple",
        "goimports",
        "gofmt",
        "gocyclo",
        "golint",
        "ineffassign"
      ],
      "Install": false,
      "Deadline": "30s",
      "Test": true
    }

### Running all checks, even if one failed

The default behaviour for `go-test-runner` is to exit if a check has failed. To
override that behaviour and run all checks, please provide the `-ignore-errors`
flag.

## Contributing

The `go-test-runner` is by no means done. It works well for most of our use
cases. But if you have other use cases you'd like it to cover, please submit an
[issue](https://github.com/agrea/go-test-runner/issues).

If you'd like to contribute by implementing more features or fixing bugs, please
don't hesitate to raise a PR. If you need any guidance, check with
[@sebdah](https://twitter.com/sebdah) or
[@AgreaIO](https://twitter.com/AgreaIO), we'd love to help get you started.

## License

MIT license
