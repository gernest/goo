# goo [![Build Status](https://travis-ci.org/gernest/goo.svg)](https://travis-ci.org/gernest/goo)

goo is a commandline tool that makes it easier to develop with golang.
It simplifies installation, configuration and use of golang tool chains.

The goo tool  wraps the golang go command, enabling you to have
full power of golang with minimal efforts.

Goo manages GOROOT, and GOPATH for you. And you don't have to worry about permissions
Infact goo recomends  not to be run with root permissions, only that your home directory
is writable.

Goo supports multiple versions of golang. Meaning you can test your code against multiple versions
of golang locally easily and secure with goo.

## Where is the go toolchain installed?

Goo install go versions tn your home directory i.e $HOME/.goo. So you can just delete the
directory if you dont want to mess with golang anymore.

# features

* manages GOPATH
* manages GOROOT
* Install/uninstall multibple go versions
* auto expand github repositories.
* forget about .bashrc, .profile and all the files you would need to edit.
* cross platform(Only tested on linux, testing on other platform is underway but you can help with that.)
'

# Installation

download latest binaries [download](/releases/latest)

Put the binary in your system PATH or just somewhere easy enough to access with your console.

And incase you have already go installed and wish to test or use goo too you can go get the project

	go get github.com/gernest/goo

## How to use.

### `goo` commands

Installing latest go version
	
	goo install latest
	
Install specific go version

	goo install 1.5
	
Get the list of available go versions for download

	goo show all
	
Get the list of installed go version

	goo show i
	

Which version of go you are currently using

	goo which go

Which GOPATH you are curretly using

	goo which gopath



### Working with your go project.

The go tool and all of its power is exposed via

	goo go
	
e.g building a project

	goo go build
	


### Help with packages from github

goo helps you work with github packages easily

Getting a project

* with go

	go get github.com/gernest/goo
	
* with goo

	goo get gernest/goo
	
Testing a project

* with go

	go test github.com/gernest/goo
	
* with goo

	goo test gernest/goo



There is much more, you can see all the usage and all the commands by running

	goo 

# Author

Geofrey Ernest <geofreyernest@live.com>


# Contributing

Fork and submit a pull request.

Enjoy