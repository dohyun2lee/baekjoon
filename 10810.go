package main

import (
	"fmt"
)

func main() {
	var N, M int
	var i, j, k int

	fmt.Scanln(&N, &M)
	var basket []int = make([]int, N)

	for a := 0; a < M; a++ {
		fmt.Scanln(&i, &j, &k)
		for b := i; b <= j; b++ {
			basket[b-1] = k
		}
	}

	for i := 0; i < N; i++ {
		fmt.Printf("%d ", basket[i])
	}
}
