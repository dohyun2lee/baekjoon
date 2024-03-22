package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var A, B, C int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A, &B, &C)

	if A+B == C {
		fmt.Println("correct!")
	} else {
		fmt.Println("wrong!")
	}
}
