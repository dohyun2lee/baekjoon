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

	for {
		if N%5 == 0 {
			cnt += N / 5
			break
		} else if (N-5)%2 == 0 {
			cnt++
			N -= 5
		} else if (N-2) >= 0 {
			cnt++
			N -= 2
		} else {
			cnt = -1
			break
		}

		if N == 0 {
			break
		}
	}

	fmt.Println(cnt)
}
