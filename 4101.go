package main

import (
	"bufio"
	"fmt"
	"os"
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
		}

		if A > B {
			fmt.Fprintln(writer, "Yes")
		} else {
			fmt.Fprintln(writer, "No")
		}
	}
}
