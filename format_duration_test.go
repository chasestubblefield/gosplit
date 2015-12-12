package main

import (
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"1s", "0:01.00"},
		{"1m", "1:00.00"},
		{"1h", "1:00:00.00"},
		{"1h1m1s", "1:01:01.00"},
		{"12h34m56.789s", "12:34:56.79"},
		{"0.999s", "0:01.00"},
		{"1m59.999s", "2:00.00"},
		{"1h59m59.999s", "2:00:00.00"},
	}
	for _, c := range cases {
		d, err := time.ParseDuration(c.in)
		if err != nil {
			t.Error(err)
		}
		got := formatDuration(d)
		if got != c.want {
			t.Errorf("formatDuration(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
