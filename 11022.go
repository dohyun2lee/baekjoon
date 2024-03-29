package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, a, b int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &t)
	defer writer.Flush()

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintf(writer, "Case #%d: %d + %d = %d\n", i+1, a, b, a+b)
	}
}