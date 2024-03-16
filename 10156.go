package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var K, N, M int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &K, &N, &M)

	sum := K * N - M

	if sum <= 0 {
		fmt.Println(0)
	} else {
		fmt.Println(sum)
	}
}