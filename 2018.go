package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, l, r, sum, ans int = 0, 1, 1, 0, 0
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	if N == 1 || N == 2 {
		fmt.Fprintln(writer, 1)
	} else {
		for l <= r && r <= N {
			if sum < N {
				r++
				sum += r
			} else {
				if sum == N {
					ans++
					l++
					sum -= l
				}
			}
		}

		fmt.Fprintln(writer, ans+1)
	}
}
