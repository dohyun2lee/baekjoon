package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, Q, a, b, cmd int
	var S string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &Q)
	fmt.Fscanln(reader, &S)

	for i := 0; i < Q; i++ {
		fmt.Fscanln(reader, &cmd, &a, &b)

		if cmd == 1 {
			fmt.Fprintln(writer, numOne(S, a, b))
		} else {
			S = numTwo(S, a, b)
		}
	}
}

func numOne(S string, a, b int) int {
	var cnt int
	S += "."

	for i := a - 1; i < b; i++ {
		if S[i] != S[i+1] {
			cnt++
		}
	}

	return cnt
}

func numTwo(S string, a, b int) string {
	var res string

	for i := 0; i < len(S); i++ {
		if i+1 >= a && i+1 <= b {
			if S[i] == 90 {
				res += "A"
			} else {
				res += string(S[i] + 1)
			}
		} else {
			res += string(S[i])
		}
	}

	return res
}
