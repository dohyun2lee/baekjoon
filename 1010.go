package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	var N, M int64
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for j := 0; j < T; j++ {
		fmt.Fscanln(reader, &N, &M)
		var ans int64 = 1
		if M - N < N {
			N = M - N
		}
		for i := M - N + 1; i <= M; i++ {
			ans *= i
			for N > 0 && ans % N == 0 {
				ans /= N
				N--
			}
		}
		fmt.Fprintln(writer, ans)
	}
}
