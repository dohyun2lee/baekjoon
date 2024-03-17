package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var A, B big.Int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A)
	fmt.Fscanln(reader, &B)

	a := new(big.Int)
	fmt.Println(a.Add(&A, &B))
	fmt.Println(a.Sub(&A, &B))
	fmt.Println(a.Mul(&A, &B))
}
