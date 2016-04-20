package file

import (
	"github.com/codegangsta/cli"
	"github.com/nathan-osman/escapefromlibc/util"

	"errors"
	"fmt"
	"os"
	"path"
)

var CpCommand = cli.Command{
	Name:      "cp",
	Usage:     "copy files and directories",
	Category:  "file",
	ArgsUsage: "SRC... DEST",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "recursive, r, R",
			Usage: "copy directories recursively",
		},
	},
	Action: func(c *cli.Context) {

		// Ensure at least two arguments were provided
		if c.NArg() < 2 {
			util.AbortWithError(errors.New("at least two arguments expected"))
		}
		var (
			srcs  = c.Args()[0 : c.NArg()-1]
			dest  = c.Args()[c.NArg()-1]
			isDir = false
		)

		// Determine the current working directory
		d, err := os.Getwd()
		if err != nil {
			util.AbortWithError(err)
		}

		// Determine if the destination is a directory
		i, err := os.Stat(dest)
		if err == nil && i.IsDir() {
			isDir = true
		}

		// Behavior differs depending on how many sources were provided
		if len(srcs) == 1 {

			// Default filename matches the source
			filename := path.Base(srcs[0])

			// If dest is a directory, use it as the destination directory;
			// otherwise use it as the filename for the destination file
			if isDir {
				d = dest
			} else {
				filename = dest
			}

			// Perform the file copy
			if err = util.Copy(srcs[0], path.Join(d, filename)); err != nil {
				util.AbortWithError(err)
			}

		} else {

			// If multiple files were provided, dest *must* be a directory
			if !isDir {
				util.AbortWithError(fmt.Errorf("%s is not a directory", dest))
			}

			// Copy each of the files to the destination directory
			for _, s := range srcs {
				if err = util.Copy(s, path.Join(dest, path.Base(s))); err != nil {
					util.AbortWithError(err)
				}
			}
		}
	},
}
