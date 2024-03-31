package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var A, B, C, cnt int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &A, &B, &C)

	if C-B <= 0 {
		fmt.Fprintln(writer, -1)
	} else {
		cnt = A / (C-B)
		fmt.Fprintln(writer, cnt+1)
	}
}
