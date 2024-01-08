package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for i:=0;i<N;i++ {
		for j := N - i - 1; j > 0; j-- {
			fmt.Print(" ")
		}
		for k := 0; k < i+1; k++ {
			fmt.Print("*")
		} 
		fmt.Println()
	}
}