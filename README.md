# goo [![Build Status](https://travis-ci.org/gernest/goo.svg)](https://travis-ci.org/gernest/goo)
`goo` is a simple version manager for Go

## Where does goo install versions of Go?
`goo` installs Go versions in `$HOME/.goo/`

# Features
* Manages `$GOPATH`
* Manages `$GOROOT`
* Allows multiple versions of Go to be installed without conflict
* Auto-expand Github repositories
* Cross platform

# Installation
Download the latest `goo` binary [here](/releases/latest) and place it in `$HOME/.goo/bin`

Add the installation path to your `$PATH`:
```Bash
$ export PATH=$PATH:$HOME/.goo/bin
```

Note: If you already have a local Go installation, you can optionally install via `go get`:
```Bash
$ go get github.com/gernest/goo
```

## Usage

### Basic
Installing latest Go version:
```Bash
$ goo install latest
```
	
Install specific Go version:
```Bash
$ goo install 1.5
```
	
Get the list of available Go versions for download:
```Bash
$ goo show all
```
	
Get the list of installed Go versions:
```Bash
$ goo show i
```

Determine which version of Go is active:
```Bash
$ goo which go
```

Determine current `$GOPATH`:
```Bash
$ goo which gopath
```

The Go tool and all of its power is exposed via:
```Bash
$ goo go
```

For example, building a project:
```Bash
$ goo go build
```

### Working with Github packages
`goo` helps you work with Github packages more easily:
```Bash
$ goo get gernest/goo
```
	
And:
```Bash
$ goo test gernest/goo
```

### Additional features
`goo` provides a lot more power than documented above — to see the full list of usage options:
```Bash
$ goo
```

# Author
Geofrey Ernest <geofreyernest@live.com>


# Contributing
Fork and submit a pull request.

Enjoy!
