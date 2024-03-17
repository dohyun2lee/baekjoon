package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, M, x1, x2, y1, y2 int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)

	var suming [1025][1025]int
	var board [1025][1025]int

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if j != N {
				fmt.Fscan(reader, &board[i][j])
			} else {
				fmt.Fscanln(reader, &board[i][j])
			}
			suming[i][j] = suming[i-1][j] + suming[i][j-1] - suming[i-1][j-1] + board[i][j]
		}
	}

	for i := 0; i < M; i++ {
		fmt.Fscanln(reader, &x1, &y1, &x2, &y2)
		fmt.Fprintln(writer, suming[x2][y2]-suming[x1-1][y2]-suming[x2][y1-1]+suming[x1-1][y1-1])
	}
}
