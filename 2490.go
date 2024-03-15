package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b, c, d int

	for i := 0; i < 3; i++ {
		fmt.Fscanln(reader, &a, &b, &c, &d)
		switch a + b + c + d {
		case 0:
			fmt.Println("D")
		case 1:
			fmt.Println("C")
		case 2:
			fmt.Println("B")
		case 3:
			fmt.Println("A")
		case 4:
			fmt.Println("E")
		}
	}
}
