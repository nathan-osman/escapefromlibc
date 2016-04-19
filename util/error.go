package util

import (
	"fmt"
	"os"
)

// AbortWithError terminates the application with the specified error message.
func AbortWithError(err error) {
	Output(fmt.Sprintf("Fatal error: %s", err))
	os.Exit(1)
}
