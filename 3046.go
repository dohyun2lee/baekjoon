package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var R1, S int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &R1, &S)

	fmt.Println((2 * S) - R1)
}
