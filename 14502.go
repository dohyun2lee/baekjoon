package main

import (
	"bufio"
	"fmt"
	"os"
)

type coordinate struct {
	x, y int
}

func main() {
	var N, M int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)
	var board [][]int = make([][]int, N)
	var column []int = make([]int, M)
	var virus, wall []coordinate

	for i := 0; i < N; i++ {
		for j:=0;j<M;j++ {
			if j != M-1 {
				fmt.Fscan(reader, &column[j])
			} else {
				fmt.Fscanln(reader, &column[j])
			}
			if column[j] == 2 {
				virus = append(virus, coordinate{i, j})
			} else if column[j] == 1 {
				wall = append(wall, coordinate{i, j})
			}
		}
		board[i] = append(board[i], column...)
	}

	for i:=0;i<N;i++ {
		for j:=0;j<M;j++ {
			
		}
	}
}
