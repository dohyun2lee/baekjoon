package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, X int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	fmt.Fscanln(reader, &N, &X)
	var slice []int = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscan(reader, &slice[i])
		if slice[i] < X {
			fmt.Fprintf(writer, "%d ", slice[i])
		}
	}
}