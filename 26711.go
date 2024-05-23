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

	fmt.Fscanln(reader, &A)
	fmt.Fscanln(reader, &B)

	fmt.Fprintln(writer, A+B)
}