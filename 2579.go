package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, score int
	var step map[int]bool = make(map[int]bool)
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var stair []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &stair[i])
	}

	
}
