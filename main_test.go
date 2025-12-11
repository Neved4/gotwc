package main

import (
	"testing"
	"time"
)

func TestFormatTimeDefault(t *testing.T) {
	loc := time.FixedZone("UTC", 0)
	now := time.Date(2024, time.July, 12, 16, 3, 52, 0, loc)

	got := formatTime(now, "%Y-%m-%dT%H:%M:%S%z")
	want := "2024-07-12T16:03:52+0000"
	if got != want {
		t.Fatalf("formatTime() = %q, want %q", got, want)
	}
}

func TestFormatTimeExtensions(t *testing.T) {
	loc := time.FixedZone("TEST", 2*60*60)
	now := time.Date(2024, time.March, 5, 14, 3, 2, 0, loc)

	tests := []struct {
		format string
		want   string
	}{
		{"%C", "20"},
		{"%D", "03/05/24"},
		{"%e", " 5"},
		{"%G", "2024"},
		{"%g", "24"},
		{"%h", "Mar"},
		{"%k", "14"},
		{"%l", " 2"},
		{"%n", "\n"},
		{"%Oe", " 5"},
		{"%R", "14:03"},
		{"%r", "02:03:02 PM"},
		{"%s", "1709640182"},
		{"%T", "14:03:02"},
		{"%t", "\t"},
		{"%u", "2"},
		{"%V", "10"},
		{"%z", "+0200"},
		{"%+", "Tue Mar  5 14:03:02 TEST 2024"},
	}

	for _, tt := range tests {
		got := formatTime(now, tt.format)
		if got != tt.want {
			t.Fatalf(
				"formatTime(%q) = %q, want %q",
				tt.format, got, tt.want,
			)
		}
	}
}
