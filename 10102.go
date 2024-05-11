package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, A, B int
	var S string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	fmt.Fscanln(reader, &S)

	for i:=0;i<len(S);i++ {
		if S[i] == 65 {
			A++
		} else {
			B++
		}
	}

	if A > B {
		fmt.Println("A")
	} else if B > A {
		fmt.Println("B")
	} else {
		fmt.Println("Tie")
	}
}