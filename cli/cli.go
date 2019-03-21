package cli

import (
	"github.com/ryanvillarreal/goCookie/core"
	"gopkg.in/urfave/cli.v1"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	flags        []cli.Flag
	proxyaddr    string
	fileLocation string
	pics         bool
	target       string
	cookie       string
	requestType  string
	delay        string
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
			Value:       "./output.txt",
			Destination: &fileLocation,
		},
		cli.StringFlag{
			Name:  "pic, P",
			Usage: "--pic <s> or -p <s> Disable picture taking.",
		},
		cli.StringFlag{
			Name:        "target,t",
			Usage:       "--target http://localhost/dir or -t http://localhost/dir",
			Destination: &target,
		},
		cli.StringFlag{
			Name:        "request,r",
			Usage:       "--request GET/POST or -p GET/POST",
			Value:       "GET",
			Destination: &requestType,
		},
		cli.StringFlag{
			Name:        "delay,d",
			Usage:       "--delay (milliseconds) or -d (milliseconds)",
			Value:       "0",
			Destination: &delay,
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
	sort.Sort(cli.FlagsByName(app.Flags))
	app.Action = noArgs

	if app.Run(os.Args) == nil {
		// have to convert delay to int before passing.
		intDelay, err := strconv.Atoi(delay)
		tmpDelay := time.Duration(intDelay)
		check(err)

		// check for url validation - will only catch minor issues.
		targetAddress, err := url.Parse(proxyaddr)
		if err != nil {
			panic("Invalid Target Address" + targetAddress.String())
		}

		// check for proxy validation.
		if proxyaddr != "" {
			proxy, err := url.Parse(proxyaddr)
			if err != nil {
				panic("Invalid Proxy URI: " + proxy.String())
			}
			core.BaseLine(target, proxyaddr, requestType, tmpDelay, fileLocation)
		} else {
			core.BaseLine(target, "", requestType, tmpDelay, fileLocation)
		}
	}
}

// if no arguments are passed to the main executable.
func noArgs(c *cli.Context) error {
	// check for no flags first. Then make sure a URL and cookie are present.
	if c.NumFlags() < 1 {
		return cli.NewExitError("Please set required flags", 0)
	} else if target == "" {
		return cli.NewExitError("URL required for operation", 0)
	} else if (strings.TrimSpace(requestType) != "GET") && (strings.TrimSpace(requestType) != "POST") {
		return cli.NewExitError("Request type must be GET or POST", 2)
	}
	//return nil so app.Run(os.Args) can start.
	return nil
}
