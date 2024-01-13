package main

import "fmt"

func main() {
	var N int

	fmt.Scanln(&N)

	if N % 2 == 0{
		fmt.Println((N-1)*(N/2))
	} else {
		fmt.Println(N*((N-1)/2))
	}
    fmt.Println("2")
}