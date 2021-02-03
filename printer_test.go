package main

import (
	"fmt"
	"testing"

	"github.com/jilleJr/flog/pkg/loglevel"
)

func TestGetSkippedLevelsSlice_Empty(t *testing.T) {
	input := map[loglevel.Level]int{}
	want := 0

	got := getSkippedLevelsSlice(input)

	if len(got) != want {
		t.Errorf("slice was not empty: %#v", got)
	}
}

func TestGetSkippedLEvelsSlice_AllInOrder(t *testing.T) {
	input := map[loglevel.Level]int{
		loglevel.Trace:       1,
		loglevel.Debug:       2,
		loglevel.Information: 3,
		loglevel.Warning:     4,
		loglevel.Error:       5,
		loglevel.Critical:    6,
		loglevel.Fatal:       7,
		loglevel.Panic:       8,
	}

	want := []string{
		"1 Trace",
		"2 Debug",
		"3 Information",
		"4 Warning",
		"5 Error",
		"6 Critical",
		"7 Fatal",
		"8 Panic",
	}

	got := getSkippedLevelsSlice(input)

	if len(got) != len(want) {
		t.Fatalf("slice was of wrong length: want %d, got %d: %#v", len(want), len(got), got)
	}

	for i, wantItem := range want {
		gotItem := got[i]
		if gotItem != wantItem {
			t.Errorf("slice was wrong at index %d: want %q, got %q", i, wantItem, gotItem)
		}
	}
}

func TestShouldIncludeLogInOutput_MinLevel(t *testing.T) {
	var testCases = []struct {
		input    loglevel.Level
		minLevel loglevel.Level
		want     bool
	}{
		{
			input:    loglevel.Fatal,
			minLevel: loglevel.Information,
			want:     true,
		},
		{
			input:    loglevel.Information,
			minLevel: loglevel.Information,
			want:     true,
		},
		{
			input:    loglevel.Debug,
			minLevel: loglevel.Information,
			want:     false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d/input/%v/minLevel/%v", i, tc.input, tc.minLevel), func(t *testing.T) {
			got := shouldIncludeLogInOutput(tc.input, LogFilter{MinLevel: tc.minLevel})
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestShouldIncludeLogInOutput_MaxLevel(t *testing.T) {
	var testCases = []struct {
		input    loglevel.Level
		maxLevel loglevel.Level
		want     bool
	}{
		{
			input:    loglevel.Fatal,
			maxLevel: loglevel.Information,
			want:     false,
		},
		{
			input:    loglevel.Information,
			maxLevel: loglevel.Information,
			want:     true,
		},
		{
			input:    loglevel.Debug,
			maxLevel: loglevel.Information,
			want:     true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d/input/%v/maxLevel/%v", i, tc.input, tc.maxLevel), func(t *testing.T) {
			got := shouldIncludeLogInOutput(tc.input, LogFilter{MaxLevel: tc.maxLevel})
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestShouldIncludeLogInOutput_BlacklistMask(t *testing.T) {
	var testCases = []struct {
		input         loglevel.Level
		blacklistMask loglevel.Level
		want          bool
	}{
		{
			input:         loglevel.Warning,
			blacklistMask: loglevel.Information | loglevel.Panic,
			want:          true,
		},
		{
			input:         loglevel.Information,
			blacklistMask: loglevel.Information | loglevel.Panic,
			want:          false,
		},
		{
			input:         loglevel.Panic,
			blacklistMask: loglevel.Information | loglevel.Panic,
			want:          false,
		},
		{
			input:         loglevel.Debug,
			blacklistMask: loglevel.Information | loglevel.Panic,
			want:          true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d/input/%v/blacklistMask/%v", i, tc.input, tc.blacklistMask), func(t *testing.T) {
			got := shouldIncludeLogInOutput(tc.input, LogFilter{BlacklistMask: tc.blacklistMask})
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestShouldIncludeLogInOutput_WhitelistMask(t *testing.T) {
	var testCases = []struct {
		input         loglevel.Level
		whitelistMask loglevel.Level
		want          bool
	}{
		{
			input:         loglevel.Warning,
			whitelistMask: loglevel.Information | loglevel.Panic,
			want:          false,
		},
		{
			input:         loglevel.Information,
			whitelistMask: loglevel.Information | loglevel.Panic,
			want:          true,
		},
		{
			input:         loglevel.Panic,
			whitelistMask: loglevel.Information | loglevel.Panic,
			want:          true,
		},
		{
			input:         loglevel.Debug,
			whitelistMask: loglevel.Information | loglevel.Panic,
			want:          false,
		},
		{
			input:         loglevel.Debug,
			whitelistMask: loglevel.Undefined,
			want:          true,
		},
		{
			input:         loglevel.Debug,
			whitelistMask: loglevel.Unknown,
			want:          false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d/input/%v/whitelistMask/%v", i, tc.input, tc.whitelistMask), func(t *testing.T) {
			got := shouldIncludeLogInOutput(tc.input, LogFilter{WhitelistMask: tc.whitelistMask})
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
