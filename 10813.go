package main

import (
	"fmt"
)

func main() {
	var N, M int
	var i, j int

	fmt.Scanln(&N, &M)
	var basket []int = make([]int, N)
	var copy []int = make([]int, N)

	for a:=0;a<N;a++ {
		basket[a] = a+1
		copy[a] = basket[a]
	}

	for a := 0; a < M; a++ {
		fmt.Scanln(&i, &j)
		k := j-1
		for b:=i;b<=j;b++ {
			if i==j {
				basket[j-1] = copy[j-1]
			} else {
				basket[b-1] = copy[k]
			}
			k--
		}
		for a := 0; a < N; a++ {
			copy[a] = basket[a]
		}
	}

	for a := 0; a < N; a++ {
		fmt.Printf("%d ", basket[a])
	}

	
}
