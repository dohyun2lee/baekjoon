package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	var way []int = make([]int, 1001)
	reader := bufio.NewReader(os.Stdin)

	way[1] = 1
	way[2] = 2
	way[3] = 3

	fmt.Fscanln(reader, &N)

	if N <= 3 {
		fmt.Println(way[N])
	} else {
		for i := 4; i <= N; i++ {
			way[i] = way[i-2] + way[i-1]
			way[i] %= 10007
		}

		fmt.Println(way[N] % 10007)
	}
}
