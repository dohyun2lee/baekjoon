package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var A, B, C, D int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A, &B)
	fmt.Fscanln(reader, &C, &D)

	if A+D <= C+B {
		fmt.Println(A+D)
	} else {
		fmt.Println(B+C)
	}
}