package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var A, B, V int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A, &B, &V)

	if V <= A {
		fmt.Println("1")
	} else {
		if (V-A) % (A-B) == 0 {
			fmt.Println((V-A)/(A-B)+1)
		} else {
			fmt.Println((V-A)/(A-B)+2)
		}
	}
	
}