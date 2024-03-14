package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var N, M big.Int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &M)

	x := new(big.Int)

	x = x.Add(&N, &M)

	fmt.Println(x)
}
