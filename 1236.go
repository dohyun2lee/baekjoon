package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	var N, M, r, l, R, L int
	var s string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &M)
	var Row [][]string = make([][]string, N)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &s)
		Row[i] = strings.Split(s, "")
	}

	for i:=0;i<N;i++ {
		l = 0
		for j:=0;j<M;j++ {
			if Row[i][j] == "X" {
				break
			} else {
				l++
			}
		}
		if l == M {
			L++
		}
	}

	for j:=0;j<M;j++ {
		r = 0
		for i:=0;i<N;i++ {
			if Row[i][j] == "X" {
				break
			} else {
				r++
			}
		}
		if r == N {
			R++
		}
	}

	if R > L {
		fmt.Println(R)
	} else {
		fmt.Println(L)
	}
}