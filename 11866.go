package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, K int
	var ans []int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &K)
	var jose []int

	for i := 1; i <= N; i++ {
		jose = append(jose, i)
	}

	var x int
	for len(jose) > 0 {
		x = (x + K - 1) % len(jose)
		ans = append(ans, jose[x])
		jose = append(jose[:x], jose[x+1:]...)
	}

	fmt.Fprintf(writer, "<")
	for i := 0; i < len(ans); i++ {
		if i < len(ans)-1 {
			fmt.Fprintf(writer, "%d, ", ans[i])
		} else {
			fmt.Fprintf(writer, "%d", ans[i])
		}
	}
	fmt.Fprintf(writer, ">")
}
