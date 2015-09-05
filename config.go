package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	errNotConfigured  = errors.New("goo: not configured yet, please run goo install latest")
	errOsNotSupported = errors.New("goo: your os is not yet supported")
	configFile        = "_goo.json"
	gooDir            = ".goo"
)

// osEnv is an interface that is used to get the os specific Path key and path separator
// Just in case there is differenct covetios regarding path separator
//
// for instance in *nix path key is PATH and path separator is ":"
type osEnv interface {

	// PathKey is the key that stores the PATH environmental variable.
	// e.g PATH for *nix
	PathKey() string

	// PathSep is the character used to separate different paths in the PATH environment
	// e.g : in *nix
	PathSep() string
}

// config is the configuration file used by goo.
// the default location for this file is in the .goo directory at the user's home directory.
//
// default name is _goo.json
type config struct {
	GOOOT     string      `json:"go_root"`
	GOPATH    string      `json:"go_path"`
	ActiveGo  string      `json:"active_go"`
	Installed []string    `json:"installed"`
	Releases  releaseInfo `json:"releases,omitempty"`
}

// update updates the configuration file, by setting GOROOT
// to point to the go version ver
func (c *config) update(ver string) error {
	c.ActiveGo = ver
	goRoot := filepath.Join(gooPath(), ver, "go")
	_, err := os.Stat(goRoot)
	if err != nil {
		return err
	}
	c.GOOOT = goRoot
	return c.save()
}

// save persists confguration file to disc.
func (c *config) save() error {
	cfgFile := filepath.Join(gooPath(), configFile)
	d, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(cfgFile, d, 0600)
}

// getConfig returns the goo configuration file.
func getConfig() (*config, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	gooPath := filepath.Join(usr.HomeDir, gooDir)
	_, err = os.Stat(gooPath)
	if os.IsNotExist(err) {
		return nil, errNotConfigured
	}
	cfgFile := filepath.Join(gooPath, configFile)
	b, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}
	cfg := &config{}
	if err = json.Unmarshal(b, cfg); err != nil {
		return nil, err
	}
	if cfg.Releases.Releases == nil {
		cfg.Releases = baseRelease
	}
	return cfg, nil
}

// sysEnv is a simple implementation of osEnv interface
type sysEnv struct {
	key, sep string
}

func (s *sysEnv) PathKey() string {
	return s.key
}
func (s *sysEnv) PathSep() string {
	return s.sep
}

func newOsEnv(key, sep string) osEnv {
	return &sysEnv{key, sep}
}

// getOsEnv returns osEnv of the host platform.
func getOsEnv() osEnv {
	switch runtime.GOOS {
	case "linux":
		return newOsEnv("PATH", ":")
	}
	return nil
}

// gooPath returns the default goo path. That is the directory where goo keeps all the
// stash
func gooPath() string {
	usr, err := user.Current()
	if err != nil {
		writeLn(err)
		os.Exit(1)
	}
	return filepath.Join(usr.HomeDir, gooDir)
}

// getDefaultConfig retrieves default configurations
func getDefaultConfig() *config {
	gp := gooPath()
	return &config{
		GOOOT:    filepath.Join(gp, "go"),
		GOPATH:   filepath.Join(gp, "gosrc"),
		Releases: baseRelease,
	}
}

// setup sets everything that is needed to run golang application.
func setup(cfg *config) error {
	env := getOsEnv()
	if env == nil {
		return errOsNotSupported
	}
	setPath(cfg, env)
	return nil

}

// setPath sets GOPATH and GOROOT, also adds the bin directories for GOPATH and GOROOT
// to the system PATH.
func setPath(cfg *config, env osEnv) {

	// check if go root and gopath exist

	_, err := os.Stat(cfg.GOOOT)
	if os.IsNotExist(err) {
		writeLn("couldn't find the go root")
		os.Exit(1)
	}
	_, err = os.Stat(cfg.GOPATH)
	if os.IsNotExist(err) {
		// we create GOPATH
		for _, v := range []string{"src", "bin", "pkg"} {
			os.MkdirAll(filepath.Join(cfg.GOPATH, v), 0755)
		}
	}
	goBin := filepath.Join(cfg.GOOOT, "bin")
	gosrcBin := filepath.Join(cfg.GOPATH, "bin")
	pathVar := os.Getenv(env.PathKey())

	if !strings.Contains(pathVar, goBin) {
		pathVar = pathVar + env.PathSep() + goBin
	}
	if !strings.Contains(pathVar, gosrcBin) {
		pathVar = pathVar + env.PathSep() + gosrcBin
	}
	os.Setenv("GOPATH", cfg.GOPATH)
	os.Setenv("GOROOT", cfg.GOOOT)
	os.Setenv(env.PathKey(), pathVar)

}
