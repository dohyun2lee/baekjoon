package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var A, B, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A, &B)

	for {
		if B == A {
			break
		} else if B < A {
			cnt = -2
			break
		}

		if B%10 == 1 {
			B = B / 10
			cnt++
		} else if B%2==0 {
			B /= 2
			cnt++
		} else {
			cnt = -2
			break
		}
	}

	fmt.Println(cnt + 1)
}
