package main

import (
	"fmt"
)

func main() {
	var slice []int = make([]int, 30)
	var a int

	for i:=0;i<28;i++ {
		fmt.Scanln(&a)
		
		slice[a-1] = 1
	}

	for i:=0;i<30;i++ {
		if slice[i] != 1 {
			fmt.Println(i+1)
		}
	}
}