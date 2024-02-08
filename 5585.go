package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	N = 1000 - N

	for N > 0 {
		if N >= 500 {
			cnt += N / 500
			N = N % 500
		} else if N >= 100 {
			cnt += N / 100
			N = N % 100
		} else if N >= 50 {
			cnt += N / 50
			N = N % 50
		} else if N >= 10 {
			cnt += N / 10
			N = N % 10
		} else if N >= 5 {
			cnt += N / 5
			N = N % 5
		} else {
			cnt += N
			N = 0
		}
	}

	fmt.Println(cnt)
}
