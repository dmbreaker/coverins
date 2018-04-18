package coverage

import (
	"io"
	"os"
	"testing"
)

type dummyTestDeps func(pat, str string) (bool, error)

func (d dummyTestDeps) MatchString(pat, str string) (bool, error)   { return false, nil }
func (d dummyTestDeps) StartCPUProfile(io.Writer) error             { return nil }
func (d dummyTestDeps) StopCPUProfile()                             {}
func (d dummyTestDeps) WriteHeapProfile(io.Writer) error            { return nil }
func (d dummyTestDeps) WriteProfileTo(string, io.Writer, int) error { return nil }
func (d dummyTestDeps) ImportPath() string                          { return "" }

// Flusher ...
type Flusher struct{}

// FlushProfiles flushes test profiles to disk. It works by build and executing
// a dummy list of 1 test. This is to ensure we execute the M.after() function
// (a function internal to the testing package) where all reporting (cpu, mem,
// cover, ... profiles) is flushed to disk.
func (*Flusher) FlushProfiles() {
	// Redirect Stdout/err temporarily so the testing code doesn't output the
	// regular:
	//   PASS
	//   coverage: 21.4% of statements
	// Thanks to this, we can test the output of the instrumented binary the same
	// way the normal binary is.
	oldstdout := os.Stdout
	oldstderr := os.Stderr
	os.Stdout, _ = os.Open(os.DevNull)
	os.Stderr, _ = os.Open(os.DevNull)

	tests := []testing.InternalTest{}
	benchmarks := []testing.InternalBenchmark{}
	examples := []testing.InternalExample{}
	var f dummyTestDeps
	dummyM := testing.MainStart(f, tests, benchmarks, examples)
	dummyM.Run()

	// restore stdout/err
	os.Stdout = oldstdout
	os.Stderr = oldstderr
}
