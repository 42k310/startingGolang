package main

import "testing"

func TestAppName(t *testing.T) {
	expect := "Zoo Apprication"
	actual := AppName()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
