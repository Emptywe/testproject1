package main

import (
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// O(n)
// 50 min
func testValidity(input string) bool {
	input = strings.TrimSuffix(input, "\n")
	mat, _ := regexp.Match(`^([\d]+-[a-zA-Z]+-?\b)+$`, []byte(input))
	return mat
}

// O(n)
// 15 min
func averageNumber(input string) float64 {
	var resi float64
	var avg float64

	re, _ := regexp.Compile(`\d+`)

	res := re.FindAllString(input, -1)
	for _, v := range res {
		t, _ := strconv.Atoi(v)
		resi += float64(t)
	}

	avg = resi / float64(len(res))
	return avg
}

// O(log n)
// 15 min
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

// O(log n)
// 15 min
func storyStats(input string) *Stats {
	var st Stats
	st.avgWordLength = make([]float64, 2)
	words := wholeStory(input)
	wordsArr := strings.Fields(words)

	var avgCount int
	var avgLen float64

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
	avgLen = float64(avgCount) / float64(len(wordsArr))

	st.avgWordLength[0] = math.Floor(avgLen)
	st.avgWordLength[1] = math.Ceil(avgLen)

	for _, v := range wordsArr {
		if len(v) == int(st.avgWordLength[0]) || len(v) == int(st.avgWordLength[1]) {
			st.avgList = append(st.avgList, v)
		}
	}
	sort.Strings(st.avgList)
	return &st
}
