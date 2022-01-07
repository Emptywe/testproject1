package main

import (
	"testing"
)

func TestStoryStats(t *testing.T) {

	seq := "1-Hello-2-World-3-Test-4-Run"
	expected := Stats{
		shortestWord:  "Run",
		longestWord:   "Hello",
		avgWordLength: 4,
		avgList:       []string{"Test"},
	}

	result := storyStats(seq)
	if result.shortestWord != expected.shortestWord || result.longestWord != expected.longestWord || result.avgList[0] != expected.avgList[0] || result.avgWordLength != expected.avgWordLength {
		t.Errorf("Incorrect story stats result. Expect %v, got %v", expected, result)
	}

}
