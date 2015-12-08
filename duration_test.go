package main

import (
	"testing"
	"time"
)

func TestDurationString(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"1s", "0:01.00"},
		{"1m", "1:00.00"},
		{"1h", "1:00:00.00"},
		{"1h1m1s", "1:01:01.00"},
	}
	for _, c := range cases {
		d, err := time.ParseDuration(c.in)
		if err != nil {
			t.Error(err)
		}
		got := duration(d).String()
		if got != c.want {
			t.Errorf("duration(%q).String() == %q, want %q", c.in, got, c.want)
		}
	}
}
