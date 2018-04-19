package main

import (
	"flag"
	"os"
	"path"
	"strings"
	"testing"
)

// TestMain is a alternate entry point
func TestMain(m *testing.M) {
	// Parse the command line using the stdlib flag package so the flags
	// defined in the testing package are populated. This include the
	// -coverprofile flag.
	flag.Parse()

	// Strip command line arguments that were for the testing package and
	// that are now handled. This will remove arguments that wouldn't be
	// recognised when using a 3rd party command line parser like
	// github.com/urfave/cli.
	var programArgs []string
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.") ||
			strings.HasPrefix(arg, "-httptest.") {
			continue
		}
		programArgs = append(programArgs, arg)
	}
	os.Args = programArgs

	// If the test binary name is $program_name or $program_name.coverage,
	// we are being asked to run the coverage-instrumented program. So call
	// main() directly.
	// panic(path.Base(os.Args[0]))
	if path.Base(os.Args[0]) == "covtest" ||
		path.Base(os.Args[0]) == "covtest.coverage" {
		main()
		return
	}

	// Run unit-tests as usual.
	os.Exit(m.Run())
}
