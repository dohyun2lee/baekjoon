package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, sum, max int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var cost []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &cost[i])
		if max < cost[i] {
			max = cost[i]
		}
		sum += cost[i]
	}

	fmt.Println(sum-max)
}
