package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, M, K, a, b, c, d, ans int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)
	var board [][]int = make([][]int, N)

	for i := 0; i < N; i++ {
		board[i] = make([]int, M)
		for j := 0; j < M; j++ {
			if j != M-1 {
				fmt.Fscan(reader, &board[i][j])
			} else {
				fmt.Fscanln(reader, &board[i][j])
			}
		}
	}

	fmt.Fscanln(reader, &K)

	for i := 0; i < K; i++ {
		ans = 0
		fmt.Fscanln(reader, &a, &b, &c, &d)

		for j:=a-1;j<c;j++ {
			for k:=b-1;k<d;k++ {
				ans += board[j][k]
			}
		}

		fmt.Fprintln(writer, ans)
	}
}
