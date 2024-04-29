package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var A, B int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		fmt.Fscanln(reader, &A, &B)

		if A == 0 && B == 0 {
			break
		} else {
			fmt.Fprintln(writer, A+B)
		}
	}
}