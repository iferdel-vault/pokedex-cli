package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		want  []string
	}{
		{input: "hello help", want: []string{"hello", "help"}},
	}

	for _, tc := range cases {
		got := cleanInput(tc.input)
		want := tc.want
		if len(got) != len(want) {
			t.Errorf("got %q, want %v", got, want)
		}
	}
}
