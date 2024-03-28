package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, blank, star int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	star = (2 * N) - 1
	blank = 0

	for i := 0; i < N; i++ {
		for j := 0; j < blank; j++ {
			fmt.Fprint(writer, " ")
		}
		for j := 0; j < star; j++ {
			fmt.Fprint(writer, "*")
		}
		blank++
		star -= 2
		fmt.Fprintln(writer)
	}
}
