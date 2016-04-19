package util

import (
	"fmt"
	"io"
	"os"
)

// Output writes the specified message to STDERR with newline appended.
func Output(msg string) {
	fmt.Fprintf(os.Stderr, "%s\n", msg)
}

// OpenInput attempts to open the specified file for input, interpreting "-" as
// STDIN.
func OpenInput(filename string) (io.Reader, error) {
	if filename == "-" {
		return os.Stdin, nil
	} else {
		return os.Open(filename)
	}
}

// OpenOutput attempts to open the specified file for output, interpreting "-"
// as STDOUT.
func OpenOutput(filename string) (io.Writer, error) {
	if filename == "-" {
		return os.Stdout, nil
	} else {
		return os.Create(filename)
	}
}
