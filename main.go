package main

import (
	"bufio"
	"fmt"
	"github.com/Emptywe/testproject1/randomExamples"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout,"Enter sequence:")
		seq, _ := reader.ReadString('\n')
		Start(seq)
		fmt.Fprint(os.Stdout, "Again? Y/N \n")
		for {
			goOn, _ := reader.ReadString('\n')
			if goOn == "N\n" || goOn == "n\n" {
				fmt.Fprint(os.Stdout, "Run random tests? Y/N\n")
				te, _ := reader.ReadString('\n')
				if te == "Y\n" || te == "y\n" {
					fmt.Fprint(os.Stdout, "Valid or invalid sequence? T/F\n")
					form, _ := reader.ReadString('\n')
					if form == "T\n" || form == "t\n" || form == "True\n" || form == "ture\n" {
						fmt.Fprint(os.Stdout, "Random valid test \n")
						str1 := randomExamples.Generate(true)
						fmt.Fprint(os.Stdout, str1, "\n")
						Start(str1)
					} else if form == "F\n" || form == "f\n" || form == "False\n" || form == "false\n" {
						fmt.Fprint(os.Stdout, "Random invalid test \n")
						str2 := randomExamples.Generate(false)
						fmt.Fprint(os.Stdout, str2, "\n")
						Start(str2)
					}
					return
				} else {
					return
				}
			} else if goOn == "Y\n" || goOn == "y\n" {
				break
			} else {
				fmt.Fprint(os.Stdout, "I don't understand, again? Y/N \n")
				continue
			}
		}
	}

}
