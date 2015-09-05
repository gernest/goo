package main

import (
	"runtime"
	"sort"

	"github.com/blang/semver"
)

// deonload is information about the go distribution archive.
type download struct {
	Os   string
	Arch string
	URL  string
}

// releases contains a list of downloads for a specific version
type release struct {
	Version   string
	Downloads []download
}

// GetDownload returns the go vversion for the host system for download.
func (r *release) GetDownload() *download {
	for _, v := range r.Downloads {
		if v.Arch == runtime.GOARCH && v.Os == runtime.GOOS {
			return &v
		}
	}
	return nil

}

// releaseInfo tracks information about golang releases
type releaseInfo struct {
	Version  string
	Releases releases
}

func (r *releaseInfo) Latest() string {
	sort.Sort(r.Releases)
	return r.Releases[0].Version
}

func (r *releaseInfo) Available() releases {
	sort.Sort(r.Releases)
	return r.Releases
}
func (r *releaseInfo) Find(ver string) *release {
	for _, v := range r.Releases {
		if v.Version == ver {
			return v
		}
	}
	return nil
}

type releases []*release

func (r releases) Len() int {
	return len(r)
}

func (r releases) Less(i, j int) bool {
	v1, err := semver.Parse(r[i].Version)
	if err != nil {
		panic(err)
	}
	v2, err := semver.Parse(r[j].Version)
	if err != nil {
		writeLn(r[j].Version)
		panic(err)
	}
	return v1.LT(v2)
}

func (r releases) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
