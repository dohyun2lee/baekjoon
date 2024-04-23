package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var L, P, count, x int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &L, &P)
	count = L * P

	for i := 0; i < 5; i++ {
		fmt.Fscan(reader, &x)
		fmt.Fprintf(writer, "%d ", x-count)
	}
}
