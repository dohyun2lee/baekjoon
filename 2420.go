package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, M int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &M)

	if N > M {
		fmt.Println(N-M)
	} else {
		fmt.Println(M-N)
	}
}