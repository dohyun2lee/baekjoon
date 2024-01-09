package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N int
	var a = 2
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for i:=1;i<=N;i++ {
		a = a * 2 -1
	}

	fmt.Println(a*a)
}