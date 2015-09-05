package main

import (
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/codegangsta/cli"
	"github.com/gernest/downloader"
	"github.com/gernest/kemi"
)

func (b *base) goCMD(ctx *cli.Context) {
	var cmd *exec.Cmd
	cmd = exec.Command("go", "help")
	if len(os.Args) > 2 {
		cmd = exec.Command("go", os.Args[2:]...)
	}
	cmd.Stdout = stdOut
	cmd.Run()
}

func (b *base) installGo(ctx *cli.Context) {
	ver := ctx.Args().First()
	rel := b.cfg.Releases.Find(ver)
	if rel == nil {
		writef("oops! we cant find releases for %s \n", ver)
		return
	}
	if down := rel.GetDownload(); down != nil {
		dFIle := downloadFile(down.URL)
		if dFIle == "" {
			writeLn("download failed")
			return
		}

		dest := filepath.Join(gooPath(), ver)
		if err := kemi.Unpack(dFIle, dest); err != nil {
			writef("somefish installing %v \n", err)
			return
		}
		writeLn("successful installed ", ver)

		err := b.cfg.update(ver)
		if err != nil {
			writeLn(err)
		}
	}
}

func downloadFile(urls string) string {
	dUrl, err := url.Parse(urls)
	if err != nil {
		writef("some fish parsing %s %v\n", urls, err)
		return ""
	}
	n := strings.Split(dUrl.Path, "/")

	tmpFile, err := ioutil.TempFile("", n[len(n)-1])
	if err != nil {
		writef("some fish opening temp file %s %v \n", dUrl.Path, err)
	}
	defer tmpFile.Close()

	fileDl, err := downloader.NewFileDl(urls, tmpFile, 0)
	if err != nil {
		writeLn(err)
		return ""
	}
	var exit = make(chan bool)
	var resume = make(chan bool)
	var pause bool
	var success bool
	var wg sync.WaitGroup
	wg.Add(1)
	fileDl.OnStart(func() {
		writef("donloading %s \n", dUrl.Path)
		format := "\033[2K\r%v/%v [%s] %v byte/s %v"
		for {
			status := fileDl.GetStatus()
			var i = float64(status.Downloaded) / float64(fileDl.Size) * 50
			h := strings.Repeat("=", int(i)) + strings.Repeat(" ", 50-int(i))

			select {
			case <-exit:
				writef(format, status.Downloaded, fileDl.Size, h, 0, "[FINISH]")
				writeLn("finside downloading")
				success = true
				wg.Done()
			default:
				if !pause {
					time.Sleep(time.Second * 1)
					writef(format, status.Downloaded, fileDl.Size, h, status.Speeds, "[DOWNLOADING]")
					os.Stdout.Sync()
				} else {
					writef(format, status.Downloaded, fileDl.Size, h, 0, "[PAUSE]")
					os.Stdout.Sync()
					<-resume
					pause = false
				}
			}
		}
	})

	fileDl.OnPause(func() {
		pause = true
	})

	fileDl.OnResume(func() {
		resume <- true
	})

	fileDl.OnFinish(func() {
		exit <- true
	})

	fileDl.OnError(func(errCode int, err error) {
		writeLn(errCode, err)
	})
	fileDl.Start()
	wg.Wait()

	if success {
		return tmpFile.Name()
	}
	return ""
}

func (b *base) useGo(ctx *cli.Context) {
	ver := ctx.Args().First()
	switch ver {
	case "":
		writeLn("please specify the version of go")
	case "latest":
		b.switchGo(b.cfg.Releases.Latest())
	default:
		b.switchGo(ver)
	}
}

func (b *base) switchGo(ver string) {
	var found bool
	for _, v := range b.cfg.Installed {
		if v == ver {
			found = true
			break
		}
	}
	if found {
		err := b.cfg.update(ver)
		if err != nil {
			writeLn(err)
			return
		}
		writeLn("switched to ", ver)
		return
	}
	writef("oops! %s is not installled yet %v\n", ver, b.cfg)
}

func (b *base) configGo(ctx *cli.Context) {}
func (b *base) whichComponent(ctx *cli.Context) {
	switch ctx.Args().First() {
	case "goroot":
		writeLn(b.cfg.GOOOT)
	case "gopath":
		writeLn(b.cfg.GOPATH)
	case "go":
		cmd := exec.Command("go", "version")
		cmd.Stdout = stdOut
		cmd.Run()
	}
}
func (b *base) listGo(ctx *cli.Context) {
	if ctx.BoolT("i") {
		writeLn("list")
	}
}
func (b *base) uninstallGo(ctx *cli.Context) {
	ver := ctx.Args().First()
	switch ver {
	case "":
		writeLn("please specify version of go you want to uninstall")
	default:
		var found bool
		for _, v := range b.cfg.Installed {
			if v == ver {
				found = true
				break
			}
		}
		if found {
			err := os.RemoveAll(filepath.Join(gooPath(), ver))
			if err != nil {
				writeLn("some fish ", err)
				return
			}
			writeLn("successful uninstalled ", ver)
			return
		}
		writeLn("can't find ", ver, " no such version of go is installed")
	}
}
