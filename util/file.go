package util

import (
	"io"
	"os"
)

// Copy copies the content of the source file to the destination file.
func Copy(src, dest string) error {

	// Open the input file
	i, err := os.Open(src)
	if err != nil {
		return err
	}
	defer i.Close()

	// Open the output file
	o, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer o.Close()

	// Copy the contents
	_, err = io.Copy(o, i)
	return err
}
