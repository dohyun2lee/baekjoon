package main

import "fmt"

func main() {
	var N int
	var s string
	
	fmt.Scanln(&N)

	for i:=0;i<N/4;i++ {
		s += "long "
	}

	fmt.Println(s + "int")
}