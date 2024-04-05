package main

import (
	"fmt"
)

func main() {
	var N, M int

	fmt.Scanln(&N, &M)

	if N == M {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
