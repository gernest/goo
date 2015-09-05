package main

import (
	"testing"
)

func TestExpand(t *testing.T) {
	src := []string{
		"bear", "clara/bear",
	}

	expected := []string{
		"bear", "github.com/clara/bear",
	}
	expand := expandRepo(src)
	for k, v := range expand {
		if v != expected[k] {
			t.Errorf("expected %s got 5s \n", expected[k], v)
		}
	}
}
