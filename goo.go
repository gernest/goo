package main

import (
	"io"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

const (
	version = "0.1.0"
	appName = "goo"
)

var stdOut io.Writer = os.Stdout

type base struct {
	*cli.App
	cfg *config
}

func getHelp(cmd string) string {
	b, err := Asset("help/" + cmd)
	if err != nil {
		return ""
	}
	return string(b)
}

func (b *base) init() {
	cfg, err := getConfig()
	if err != nil {
		writeLn(err)
		writeLn("trying to use defaults instead")
		cfg = getDefaultConfig()
	}
	b.cfg = cfg
	err = setup(cfg)
	if err != nil {
		writeLn(err)
		os.Exit(1)
	}
	b.Action = func(ctx *cli.Context) {
		switch ctx.Args().First() {
		case "":
			cmd := exec.Command(os.Args[0], "-h")
			cmd.Stdout = stdOut
			cmd.Run()
		default:
			cmd := exec.Command("go", "help")
			if len(os.Args) > 1 {
				cmd = exec.Command("go", expandRepo(os.Args[1:])...)
			}
			cmd.Stdout = stdOut
			cmd.Run()
		}
	}
	b.Commands = []cli.Command{
		cli.Command{
			Name:        "use",
			Usage:       "Switches go versions",
			Description: getHelp("use"),
			Action:      b.useGo,
		},
		cli.Command{
			Name:            "go",
			Usage:           "Runs the go command",
			Action:          b.goCMD,
			SkipFlagParsing: true,
		},
		cli.Command{
			Name:   "which",
			Usage:  "Answers many questions about goo and go",
			Action: b.whichComponent,
		},
		cli.Command{
			Name:   "install",
			Usage:  "Installs a speific version of go",
			Action: b.installGo,
		},
		cli.Command{
			Name:  "show",
			Usage: "Lists go versions",
			Subcommands: []cli.Command{
				cli.Command{
					Name:      "installed",
					ShortName: "i",
					Usage:     "list installed go versions",
					Action: func(ctx *cli.Context) {
						if len(b.cfg.Installed) == 0 {
							writeLn("no go version found")
							return
						}
						for _, v := range b.cfg.Installed {
							writeLn(v)
						}
					},
				},
				cli.Command{
					Name:      "all",
					ShortName: "a",
					Usage:     "list all go versions",
					Action: func(ctx *cli.Context) {
						av := cfg.Releases.Available()
						if len(av) == 0 {
							writeLn("no go version found")
							return
						}
						for _, v := range av {
							writeLn(v.Version)
						}
					},
				},
			},
		},
		cli.Command{
			Name:   "uninstall",
			Usage:  "uninstall go versions",
			Action: b.uninstallGo,
		},
	}
}

func newApp() *base {
	b := &base{App: cli.NewApp()}
	b.Name = appName
	b.Usage = "Win at Go(Golang)"
	b.Version = version
	b.Authors = []cli.Author{
		cli.Author{
			Name:  "Geofrey Ernest",
			Email: "geofreyernest@live.com",
		},
	}
	b.init()
	return b
}

func setOutput(out io.Writer) {
	stdOut = out
}

func main() {
	app := newApp()
	app.Run(os.Args)
}
