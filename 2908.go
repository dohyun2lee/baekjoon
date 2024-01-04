package main

import (
	"fmt"
)

func main() {
	var A, B int
	
	fmt.Scanln(&A, &B)

	sumA := A/100 + (A/10)%10*10 + (A%10)*100
	sumB := B/100 + (B/10)%10*10 + (B%10)*100

	if sumA < sumB {
		fmt.Println(sumB)
	} else {
		fmt.Println(sumA)
	}
}