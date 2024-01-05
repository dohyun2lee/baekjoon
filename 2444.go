package main

import (
	"fmt"
)

func main() {	
	var N int

	fmt.Scanln(&N)

	for i:=0;i<N;i++ {
		for j:=0;j<(N-1-i);j++ {
			fmt.Print(" ")
		}
		for k:=0;k<(2*i+1);k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i:=N-1;i>0;i-- {
		for j:=0;j<N-i;j++ {
			fmt.Print(" ")
		}
		for k:=0;k<(2*i-1);k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}