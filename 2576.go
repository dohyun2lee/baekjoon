package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var sum, min, x int = 0, 100, 0
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < 7; i++ {
		fmt.Fscanln(reader, &x)
		if x%2 != 0 {
			sum += x
			if x < min {
				min = x
			}
		}
	}

	if sum == 0 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, sum)
		fmt.Fprintln(writer, min)
	}
}
