package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, M, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &M)

	if N == 1 {
		cnt = 1
	} else if N == 2 {
		if 4 < (M+1)/2 {
			cnt = 4
		} else {
			cnt = (M + 1) / 2
		}
	} else if M < 7 {
		if 4 < M {
			cnt = 4
		} else {
			cnt = M
		}
	} else {
		cnt = M - 2
	}

	fmt.Println(cnt)
}
