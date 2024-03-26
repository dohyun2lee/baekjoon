package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var x, sum int
	reader := bufio.NewReader(os.Stdin)

	for i:=0;i<5;i++ {
		fmt.Fscanln(reader, &x)
		sum += x
	}

	fmt.Println(sum)
}