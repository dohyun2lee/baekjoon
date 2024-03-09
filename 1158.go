package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, K int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &K)
	var Josephus []int = make([]int, N)
	var Ans []int

	for i := 0; i < N; i++ {
		Josephus[i] = i + 1
	}

	index := 0
	for len(Josephus) > 0 {
		index = (index + K - 1) % len(Josephus)
		Ans = append(Ans, Josephus[index])
		Josephus = append(Josephus[:index], Josephus[index+1:]...)
	}

	for i := 0; i < len(Ans); i++ {
		if i == 0 {
			fmt.Fprintf(writer, "<%d", Ans[i])
		} else {
			fmt.Fprintf(writer, ", %d", Ans[i])
		}
	}
	fmt.Fprintln(writer, ">")
}
