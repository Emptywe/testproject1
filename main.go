package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter sequence:")
		seq, _ := reader.ReadString('\n')
		Start(seq)
		fmt.Fprint(os.Stdout, "Again? Y/N \n")
		for {
			goOn, _ := reader.ReadString('\n')
			if goOn == "N\n" || goOn == "n\n" {
				return
			} else if goOn == "Y\n" || goOn == "y\n" {
				break
			} else {
				fmt.Fprint(os.Stdout, "I don't understand, again? Y/N \n")
				continue
			}
		}
	}
}
