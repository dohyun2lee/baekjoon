package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	fir := N / 100 * 78
	sec := (N / 10 * 8) + (N / 10 * 2 / 100 * 78)

	fmt.Fprintln(writer, fir, sec)
}