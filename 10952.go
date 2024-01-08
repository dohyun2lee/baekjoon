package main

import "fmt"

func main() {
	var A, B int

	for  {
		fmt.Scanln(&A, &B)
		if A == 0 || B == 0 {
			break
		}
		fmt.Println(A + B)
	}
}