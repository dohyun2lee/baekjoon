package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, a, b, c, d, Mx, My, mx, my int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var ans map[int]int = make(map[int]int)
	var board [1001][1001]int

	for i := 1; i <= N; i++ {
		fmt.Fscanln(reader, &a, &b, &c, &d)

		Mx, mx = max(a, c)
		My, my = max(b, d)

		for j:=mx;j<=Mx;j++ {
			for k:=my;k<My;k++ {
				board[j][k] = i
			}
		}
	}

	for i:=0;i<1001;i++ {
		for j:=0;j<1001;j++ {
			if board[i][j] != 0 {
				ans[board[i][j]]++
			}
		}
	}

	for i:=1;i<=N;i++ {
		fmt.Fprintln(writer, ans[i])
	}
}

func max(a, b int) (M, m int) {
	if a > b {
		M = a
		m = b
	} else {
		M = b
		m = a
	}

	return M, m
}
