package main

import (
	"regexp"
	"strconv"
	"strings"
)

func testValidity(input string) bool {
	input = strings.TrimSuffix(input, "\n")
	mat, _ := regexp.Match(`^([\d]+-[a-zA-Z]+-?\b)+$`, []byte(input))
	return mat
}

func averageNumber(input string) int {
	var resi int
	var avg int

	re, _ := regexp.Compile(`\d+`)

	res := re.FindAllString(input, -1)
	for _, v := range res {
		t, _ := strconv.Atoi(v)
		resi += t
	}

	avg = resi / len(res)
	return avg
}

func wholeStory(input string) string {
	return ""
}

func storyStats(input string) *Stats {
	return nil
}
