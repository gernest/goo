package main

// this is the base information about golang versions supported by goo
var baseRelease = releaseInfo{
	Version: "0.1.0",
	Releases: releases{
		{
			Version: "1.5.0",
			Downloads: []download{
				{
					"linux",
					"amd64",
					"https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz",
				},
				{
					"linux",
					"386",
					"https://storage.googleapis.com/golang/go1.5.linux-386.tar.gz",
				},
			},
		},
		{
			Version: "1.4.2",
			Downloads: []download{
				{
					"linux",
					"amd64",
					"https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz",
				},
				{
					"linux",
					"386",
					"https://storage.googleapis.com/golang/go1.4.2.linux-386.tar.gz",
				},
			},
		},
	},
}
