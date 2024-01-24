package main

import (
	"fmt"
)

func main() {
	var N int
	var ans int = 1

	fmt.Scanln(&N)

	for i:=1;i<=N;i++ {
		ans *= i
	}

	fmt.Println(ans)
}