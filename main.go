package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Fprint(os.Stdout,"Enter sequence:")
	seq, _ := reader.ReadString('\n')
	fmt.Fprint(os.Stdout, seq)
	valid := testValidity(seq)
	if valid {
		fmt.Fprint(os.Stdout, "Validation: true\n")
		fmt.Fprint(os.Stdout, "Average number: ", averageNumber(seq), "\n")
		fmt.Fprint(os.Stdout, "Words: ", wholeStory(seq), "\n")
		stat := storyStats(seq)
		fmt.Fprintf(os.Stdout, "Story Stats:\nShortest word: %s\nLongest word: %s\n"+
			"Averege letter in word number: %d\nAverege words:\n%s\n", stat.shortestWord, stat.longestWord, stat.avgWordLength, strings.Join(stat.avgList, "\n"))
	} else {
		fmt.Fprint(os.Stdout, "Validation: false\n")
	}
}
