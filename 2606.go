package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, K, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var com map[int]bool = make(map[int]bool)
	fmt.Fscanln(reader, &K)
	var con [][]int = make([][]int, K)

	com[1] = true

	for i := 0; i < K; i++ {
		con[i] = make([]int, 2)
		fmt.Fscanln(reader, &con[i][0], &con[i][1])
	}

	sort.Slice(con, func(i, j int) bool {
		if con[i][0] == con[j][0] {
			return con[i][1] < con[j][1]
		}
		return con[i][0] < con[j][0]
	})

	for i := 0; i < K; i++ {
		if com[con[i][0]] {
			com[con[i][1]] = true
		} else if com[con[i][1]] {
			com[con[i][0]] = true
 		}
	}
	for i := 0; i < K; i++ {
		if com[con[i][0]] {
			com[con[i][1]] = true
		} else if com[con[i][1]] {
			com[con[i][0]] = true
 		}
	}

	for _, val := range com {
		if val {
			cnt++
		}
	}

	fmt.Println(cnt - 1)
}
