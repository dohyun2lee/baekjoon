package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var A, B string
	var ans int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A)
	fmt.Fscanln(reader, &B)

	for len(B) >= len(A){
		if B[len(B)-1] == 65 {
			B = B[:len(B)-1]
		} else {
			B = B[:len(B)-1]
			B = Reverse(B)
		}

		if A == B {
			ans = 1
			break
		}
	}

	fmt.Println(ans)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
