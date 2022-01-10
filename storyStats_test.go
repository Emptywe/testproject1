package main

import (
	"testing"
)

func TestStoryStats(t *testing.T) {

	seq := "1-Hello-2-World-3-Test-4-Run"
	expected := Stats{
		shortestWord:  "Run",
		longestWord:   "Hello",
		avgWordLength: []float64{4,5},
		avgList:       []string{"Hello","Test","World"},
	}

	result := storyStats(seq)
	if result.shortestWord != expected.shortestWord || result.longestWord != expected.longestWord || result.avgList[0] != expected.avgList[0] || result.avgWordLength[0] != expected.avgWordLength[0] || result.avgWordLength[1] != expected.avgWordLength[1]{
		t.Errorf("Incorrect story stats result. Expect %v, got %v", expected, result)
	}

}
