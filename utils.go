package main

import (
	"fmt"
	"strings"
)

func write(a ...interface{}) {
	fmt.Fprint(stdOut, a...)
}

func writeLn(a ...interface{}) {
	fmt.Fprintln(stdOut, a...)
}

func writef(format string, a ...interface{}) {
	fmt.Fprintf(stdOut, format, a...)
}

func expandRepo(src []string) []string {
	githubPrefix := "github.com/"
	for k, v := range src {
		if strings.Contains(v, "/") {
			sec := strings.Split(v, "/")
			if len(sec) == 2 {
				src[k] = githubPrefix + strings.TrimPrefix(v, "/")
			}
		}
	}
	return src
}
