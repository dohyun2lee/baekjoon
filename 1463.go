package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var DP []int = make([]int, N+1)
	
	for i := 2; i <= N; i++ {
		DP[i] = DP[i-1]

		if i%2 == 0 {
			DP[i] = Min(DP[i], DP[i/2])
		}
		if i%3 == 0 {
			DP[i] = Min(DP[i], DP[i/3])
		}

		DP[i]++
	}

	fmt.Println(DP[N])
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

