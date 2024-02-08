package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var A, B, C, T int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &T)

	if T%10 != 0 {
		fmt.Println(-1)
	} else {
		A = T / 300
		B = (T % 300) / 60
		C = ((T % 300) % 60) / 10
		fmt.Println(A, B, C)
	}
}
