package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var N, M ,K int
	var matrix_one [101][101]int
	var matrix_two [101][101]int
	var matrix_ans [101][101]int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)

	for i:=0;i<N;i++ {
		for j:=0;j<M;j++ {
			if j == M-1 {
				fmt.Fscanln(reader, &matrix_one[i][j])
			} else {
				fmt.Fscan(reader, &matrix_one[i][j])
			}
		}
	}

	fmt.Fscanln(reader, &M, &K)

	for i:=0;i<M;i++ {
		for j:=0;j<K;j++ {
			if j == K-1 {
				fmt.Fscanln(reader, &matrix_two[i][j])
			} else {
				fmt.Fscan(reader, &matrix_two[i][j])
			}
		}
	}

	for i:=0;i<N;i++ {
		for j:=0;j<K;j++ {
			for k:=0;k<M;k++ {
				matrix_ans[i][j] += matrix_one[i][k] * matrix_two[k][j]
			}
		}
	}

	for i:=0;i<N;i++ {
		for j:=0;j<K;j++ {
			if j == K-1 {
				fmt.Fprintln(writer, matrix_ans[i][j])
			} else {
				fmt.Fprintf(writer, "%d ", matrix_ans[i][j])
			}
		}
	}
}