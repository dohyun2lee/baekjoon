package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, M, i, j int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)
	var num []int = make([]int, N+1)

	num[0] = 0

	for i := 1; i <= N; i++ {
		if i == N {
			fmt.Fscanln(reader, &num[i])
		} else {
			fmt.Fscanf(reader, "%d ", &num[i])
		}
		num[i] += num[i-1]
	}

	for l := 0; l < M; l++ {
		fmt.Fscanln(reader, &i, &j)
		fmt.Fprintln(writer, num[j]-num[i-1])
	}
}
