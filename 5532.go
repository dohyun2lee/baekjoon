package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var L, A, B, a, b, aMin, bMin int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &L)
	fmt.Fscanln(reader, &A)
	fmt.Fscanln(reader, &B)
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	if A % a != 0 {
		aMin = A/a + 1
	} else {
		aMin = A/a
	}

	if B % b != 0 {
		bMin = B/b + 1
	} else {
		bMin = B/b
	}

	if aMin > bMin {
		fmt.Println(L-aMin)
	} else {
		fmt.Println(L-bMin)
	}
}