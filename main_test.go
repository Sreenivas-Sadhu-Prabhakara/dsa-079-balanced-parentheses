package main

import "testing"

func TestIsBalanced(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"{[()]}", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"(", false},
		{")", false},
		{"", true},
		{"a(b)c", true}, // non-bracket characters are ignored
		{"]]", false},
	}
	for _, c := range cases {
		if got := IsBalanced(c.in); got != c.want {
			t.Errorf("IsBalanced(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
