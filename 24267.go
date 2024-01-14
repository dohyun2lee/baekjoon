package main

import "fmt"

func main() {
	var N int

	fmt.Scanln(&N)
	fmt.Println(N*(N-1)*(N-2)/6)
	fmt.Println("3")
}