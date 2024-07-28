package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var N int
	var A, B big.Int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &A, &B)

		a := new(big.Int)
		fmt.Fprintln(writer, a.Add(&A, &B))
	}
}
