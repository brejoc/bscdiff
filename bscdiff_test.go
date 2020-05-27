package main

import (
	"bytes"
	"testing"
)

func TestPrettyPrintMissingBscs(t *testing.T) {
	var res1 []searchResult

	res1 = append(res1,
		searchResult{
			line:  1,
			match: []string{"bsc#123"},
			text:  "Line in text with bsc#123",
		},
	)
	res1 = append(res1,
		searchResult{
			line:  2,
			match: []string{"bsc#321"},
			text:  "Line in text with bsc#321",
		},
	)

	buffer := &bytes.Buffer{}
	prettyPrintMissingBscs(res1, []string{"bsc#321"}, buffer)
	got := buffer.String()
	want := "2: bsc#321 -> Line in text with bsc#321\n"
	if got != want {
		t.Errorf("got \"%s\", wanted \"%s\"", got, want)
	}
}

func TestFindMissingBsc(t *testing.T) {
	var change1 []searchResult
	var change2 []searchResult

	change1 = append(change1,
		searchResult{
			line:  1,
			match: []string{"bsc#123"},
			text:  "Line in text with bsc#123",
		},
	)
	change1 = append(change1,
		searchResult{
			line:  2,
			match: []string{"bsc#321"},
			text:  "Line in text with bsc#321",
		},
	)

	change2 = append(change2,
		searchResult{
			line:  1,
			match: []string{"bsc#123"},
			text:  "Line in text with bsc#123",
		},
	)
	change2 = append(change2,
		searchResult{
			line:  2,
			match: []string{"bsc#666"},
			text:  "Line in text with bsc#666",
		},
	)
	got := findMissingBsc(change1, change2)
	want := []string{"bsc#321"}
	if got[0] != want[0] {
		t.Errorf("Got \"%s\", wanted \"%s\"", got, want)
	}
}
