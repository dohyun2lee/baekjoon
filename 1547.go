package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, ball, a, b int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	ball = 1

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &a, &b)

		if ball == a {
			ball = b
		} else if ball == b {
			ball = a
		}
	}

	fmt.Println(ball)
}
