package main

import (
	"testing"
)

func TestName(t *testing.T) {
	var tf TimeFrame
	tf.Name = "Night"
	got := Name(tf)
	want := "Night"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestDurationInSeconds(t *testing.T) {
	var tf TimeFrame
	tf.StartRealTimeIn24HourClock = 20
	tf.EndRealTimeIn24HourClock = 6
	got := DurationInSeconds(tf)
	want := 36000
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDawn(t *testing.T) {
	tests := []struct {
		input int
		want bool
	}{
		{0, false },
		{1, true },
		{300, true },
		{301, false },
	}

	for _, tc := range tests {
		got := Dawn(tc.input)
		if got != tc.want {
			t.Errorf("using %d got %t want %t", tc.input, got, tc.want)
		}
	}
}

func TestDisplayTime(t *testing.T) {
	tests := []struct {
		input int
		want string
	}{
		{301, "." },
		{300, "5:00" },
		{299, "." },
		{90, "." },
		{61, "." },
		{60, "1:00" },
		{59, "." },
		{30, "0:30" },
		{0, "DONE!!!" },
	}

	for _, tc := range tests {
		got := DisplayTime(tc.input)
		if got != tc.want {
			t.Errorf("got %s want %s", got, tc.want)
		}
	}
}
