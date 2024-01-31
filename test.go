package main

import (
	"fmt"
	"math"
)

func main(){
	var N, M float64

	fmt.Scanln(&N, &M)

	fmt.Println(math.Sqrt(math.Pow(N - M, 2)))
}