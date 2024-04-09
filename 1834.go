package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, sum int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	sum = ((N * (N+1))/2 - N) * (N+1)

	fmt.Fprintln(writer, sum)
}