package main

import (
	"bufio"
	"os"
	"fmt"
	"math/big"
)

func main() {
	var A, B big.Int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &A)
	fmt.Fscanln(reader, &B)

	a := new(big.Int)
	fmt.Fprintln(writer, a.Add(&A, &B))
}