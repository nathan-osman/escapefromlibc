package net

import (
	"github.com/codegangsta/cli"
	"github.com/nathan-osman/escapefromlibc/util"
	"gopkg.in/cheggaaa/pb.v1"

	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

var WgetCommand = cli.Command{
	Name:      "wget",
	Usage:     "retrieve remote documents via HTTP",
	Category:  "net",
	ArgsUsage: "URI",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "method",
			Value: "GET",
			Usage: "HTTP method",
		},
		cli.StringFlag{
			Name:  "O",
			Usage: "output file ('-' for STDIN)",
		},
		cli.BoolFlag{
			Name:  "quiet, q",
			Usage: "supress output",
		},
	},
	Action: func(c *cli.Context) {

		// Ensure a single argument (the URI) was provided
		if c.NArg() != 1 {
			util.AbortWithError(errors.New("single argument expected"))
		}
		rawUri := c.Args()[0]

		// Determine the correct name for the output file and open it
		outFilename := c.String("O")
		if outFilename == "" {
			u, err := url.Parse(rawUri)
			if err != nil {
				util.AbortWithError(err)
			}
			if u.Path == "" {
				outFilename = "index.html"
			} else {
				outFilename = path.Base(u.Path)
			}
		}
		out, err := util.OpenOutput(outFilename)
		if err != nil {
			util.AbortWithError(err)
		}

		// Build the request
		req, err := http.NewRequest(c.String("method"), rawUri, nil)
		if err != nil {
			util.AbortWithError(err)
		}

		// Create and send the request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			util.AbortWithError(err)
		}

		// In keeping with wget behavior, check the response code
		if resp.StatusCode >= 400 {
			util.AbortWithError(errors.New(resp.Status))
		}

		in := io.Reader(resp.Body)

		// Show download progress if not surpressed
		if !c.Bool("quiet") {
			bar := pb.New64(resp.ContentLength).SetUnits(pb.U_BYTES)
			bar.Output = os.Stderr
			bar.SetRefreshRate(500 * time.Millisecond)
			bar.ShowSpeed = true
			bar.Start()
			in = bar.NewProxyReader(in)
		}

		// Download the response body, copying it to the output file
		_, err = io.Copy(out, in)
		if err != nil {
			util.AbortWithError(err)
		}

		// Be tidy
		resp.Body.Close()
	},
}
