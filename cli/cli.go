package cli

import (
	"goCookie/core"
	"gopkg.in/urfave/cli.v1"
	"os"
	"sort"
	"strings"
)

var (
	flags        []cli.Flag
	proxyaddr    string
	fileLocation string
	pics         bool
	url          string
	cookie       string
	requestType  string
)

func init() {
	// define the app commands. - shown in order they are specified.
	flags = []cli.Flag{
		cli.StringFlag{
			Name:        "proxy, p",
			Usage:       "--proxy <127.0.0.1:8080> or -f <127.0.0.1:8080> ",
			Destination: &proxyaddr,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "--output <output> or -o <output> ",
			Destination: &fileLocation,
		},
		cli.StringFlag{
			Name:  "pic, P",
			Usage: "--pic <s> or -p <s> Disable picture taking.",
		},
		cli.StringFlag{
			Name:        "url,u",
			Usage:       "--url http://localhost/dir or -u http://localhost/dir",
			Destination: &url,
		},
		cli.StringFlag{
			Name:        "request,r",
			Usage:       "--request GET/POST or -p GET/POST",
			Value:       "GET",
			Destination: &requestType,
		},
	}
}

// main menu help function.  setups all cli arguments needed.  Information provided here.
func MenuHelp() {
	app := cli.NewApp()
	app.Name = "goCookie"
	app.Version = "0.0.1"
	app.Author = "l33tllama"
	app.Usage = "Usage: goCookie.exe https://www.example.com"
	app.Flags = flags
	app.Action = noArgs
	sort.Sort(cli.FlagsByName(app.Flags))
	if app.Run(os.Args) == nil {
		if proxyaddr != "" {
			core.BaseLine(url, proxyaddr, requestType)
		} else {
			core.BaseLine(url, "", requestType)
		}
	}
}

// if no arguments are passed to the main executable.
func noArgs(c *cli.Context) error {
	// check for no flags first. Then make sure a URL and cookie are present.
	if c.NumFlags() < 1 {
		return cli.NewExitError("Please set required flags", 2)
	} else if url == "" {
		return cli.NewExitError("URL required for operation", 2)
	} else if strings.TrimSpace(requestType) != "GET" || strings.TrimSpace(requestType) == "POST" {
		return cli.NewExitError("Request type must be GET or POST", 2)
	}
	return nil
}
