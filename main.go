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

}
