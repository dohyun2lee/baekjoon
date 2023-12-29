package main

import (
	"fmt"
)

func main() {
	var X, N, a, b, Sum int

	fmt.Scanln(&X)
	fmt.Scanln(&N)

	for i:=0;i<N;i++ {
		fmt.Scanln(&a, &b)
		Sum += (a * b)
	}

	if X == Sum {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}