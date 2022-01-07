package main

import (
	"regexp"
	"strings"
)

func testValidity(input string) bool {
	input = strings.TrimSuffix(input, "\n")
	mat, _ := regexp.Match(`^([\d]+-[a-zA-Z]+-?\b)+$`, []byte(input))
	return mat
}

func averageNumber(input string) int {
	return 0
}

func wholeStory(input string) string {
	return ""
}

func storyStats(input string) *Stats {
	return nil
}
