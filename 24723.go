package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	ans := 1

	for i := 0; i < N; i++ {
		ans *= 2
	}

	fmt.Fprintln(writer, ans)
}
