package main

import (
	// "container/list"
	"fmt"
	"bufio"
	"os"
)

func main() {
	var N, x int
	var num []int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for i:=0;i<N;i++ {
		fmt.Fscanf(reader, "%d ", &x)
		num = append(num, x)
	}

	fmt.Println(num)
}
