package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var X, a, b int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &X)

	for i:=1;i<=X;i++ {
		a += i
		if a >= X {
			b = i
			break
		}
	}

	if b % 2 == 0 {
		fmt.Printf("%d/%d", (b-a+X),(a-X+1))
	} else {
		fmt.Printf("%d/%d", (a-X+1), (b-a+X))
	}
}