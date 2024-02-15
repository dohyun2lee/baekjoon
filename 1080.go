package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N, M, cnt int
	var x string
	var matrix [][]string
	var matrix_ans [][]string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &M)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &x)
		X := strings.Split(x, "")
		matrix = append(matrix, X)
	}

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &x)
		X := strings.Split(x, "")
		matrix_ans = append(matrix_ans, X)
	}

	if N < 3 || M < 3 {
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				if matrix[i][j] != matrix_ans[i][j] {
					cnt = -1
					break
				} else {
					continue
				}
			}
		}
	} else {
		for i := 0; i < N-2; i++ {
			for j := 0; j < M-2; j++ {
				if matrix[i][j] != matrix_ans[i][j] {
					cnt++
					for x := i; x < i+3; x++ {
						for y := j; y < j+3; y++ {
							if matrix[x][y] == "0" {
								matrix[x][y] = "1"
							} else {
								matrix[x][y] = "0"
							}
						}
					}
				} else {
					continue
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if matrix[i][j] != matrix_ans[i][j] {
				cnt = -1
				break
			} else {
				continue
			}
		}
	}

	fmt.Println(cnt)
}
