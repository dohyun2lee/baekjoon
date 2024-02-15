package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var A, B, C, M, work, cnt, fatigue int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A, &B, &C, &M)

	for cnt < 24 {
		if fatigue+A > M {
			cnt++
			fatigue -= C
		} else {
			cnt++
			work += B
			fatigue += A
		}
		if fatigue <= 0 {
			fatigue = 0
		}
	}

	fmt.Println(work)
}
