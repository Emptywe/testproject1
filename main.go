package main

import (
	"bufio"
	"fmt"
	"os"
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
	} else {
		fmt.Fprint(os.Stdout, "Validation: false\n")
	}
}
