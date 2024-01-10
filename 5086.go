package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var a, b int
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fscanln(reader, &a, &b)
		
		if a == 0 && b == 0 {
			break
		} 
		
		if a%b == 0 {
			fmt.Println("multiple")
		} else if b%a == 0 {
			fmt.Println("factor")
		} else {
			fmt.Println("neither")
		}
	}
}