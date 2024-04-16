package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, rice, max int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		max = 0
		for j := 0; j < 5; j++ {
			if j != 4 {
				fmt.Fscan(reader, &rice)
			} else {
				fmt.Fscanln(reader, &rice)
			}

			if rice > max {
				max = rice
			}
		}

		fmt.Fprintf(writer, "Case #%d: %d\n", i+1, max)
	}
}
