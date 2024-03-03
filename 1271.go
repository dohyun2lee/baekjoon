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
	y := new(big.Int)

	x = x.Div(&N, &M)
	y = y.Mod(&N, &M)

	fmt.Println(x)
	fmt.Println(y)
}
