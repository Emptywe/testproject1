package main

import (
	"regexp"
	"sort"
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
	var words string

	re, _ := regexp.Compile(`[a-zA-Z]+`)

	res := re.FindAllString(input, -1)
	for i, v := range res {
		if i == len(res)-1 {
			words += v
			break
		}
		words += v + " "
	}
	return words
}

func storyStats(input string) *Stats {
	var st Stats
	words := wholeStory(input)
	wordsArr := strings.Fields(words)

	var avgCount int

	for i, v := range wordsArr {
		if i == 0 {
			st.shortestWord = v
			st.longestWord = v
			avgCount += len(v)
			continue
		}
		avgCount += len(v)
		if len(v) == len(wordsArr[i-1]) {
			continue
		}
		if len(v) > len(st.longestWord) {
			st.longestWord = v
		} else if len(v) < len(st.shortestWord) {
			st.shortestWord = v
		}

	}
	st.avgWordLength = avgCount / len(wordsArr)

	for _, v := range wordsArr {
		if len(v) == st.avgWordLength {
			st.avgList = append(st.avgList, v)
		}
	}
	sort.Strings(st.avgList)
	return &st
}
