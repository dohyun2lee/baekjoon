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
	var pibonacchi []int = make([]int, N+1)

	pibonacchi[0] = 0
	pibonacchi[1] = 1

	if N < 2 {
		fmt.Println(pibonacchi[N])
	} else {
		for i := 2; i <= N; i++ {
			pibonacchi[i] = pibonacchi[i-2] + pibonacchi[i-1]
		}

		fmt.Println(pibonacchi[N])
	}
}
