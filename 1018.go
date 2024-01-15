package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var N, M, ans int
	var cnt int = 64
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &M)
	var board = [][]string{} 

	for i:=0;i<N;i++ {
		var a string
		fmt.Fscanln(reader, &a)
		A := strings.Split(a, "")
		board = append(board, A)
	}

	for i:=0;i<N-7;i++ {
		for j:=0;j<M-7;j++ {
			var cnt1 int
			var color string
			color = board[i][j]
			for k:=i;k<i+8;k++ {
				for l:=j;l<j+8;j++ {

					if (i+j)%2 == 0 {
						if (k+l)%2 == 0 {
							if board[k][l] != color {
								cnt1++
							}
						} else {
							if board[k][j] == color {
								cnt1++
							}
						}
					} else {
						if (k+l)%2 != 0 {
							if board[k][l] != color {
								cnt1++
							}
						} else {
							if board[k][j] == color {
								cnt1++
							}
						}
					}

					if cnt1 <= cnt {
						ans = cnt1
						cnt = cnt1
					}
				}
			}
		}
	}

	fmt.Println(ans)
}