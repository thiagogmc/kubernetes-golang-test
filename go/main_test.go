package main

import "testing"

func TestGreeting(t *testing.T) {
	got := greeting("Code.education Rocks!")
	if got != "<b>Code.education Rocks!</b>" {
		t.Errorf("greeting return = %s; want: Code.education Rocks!", got)
	}
}