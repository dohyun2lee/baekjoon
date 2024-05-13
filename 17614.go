package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdio = bufio.NewReadWriter(
	bufio.NewReader(os.Stdin),
	bufio.NewWriter(os.Stdout),
)

func getClap(n int) int {
	cnt := 0
	for n > 0 {
		digit := n % 10
		n /= 10
		if digit == 3 || digit == 6 || digit == 9 {
			cnt++
		}
	}
	return cnt
}

func main() {
	defer stdio.Flush()

	var N int
	fmt.Fscan(stdio, &N)

	cnt := 0
	for i := 1; i <= N; i++ {
		cnt += getClap(i)
	}
	fmt.Fprintln(stdio, cnt)
}
