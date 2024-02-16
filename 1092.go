package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, M int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var crane []int = make([]int, N)

	for i := 0; i < N; i++ {
		if i == N-1 {
			fmt.Fscanln(reader, &crane[i])
		} else {
			fmt.Fscan(reader, &crane[i])
		}
	}

	sort.Ints(crane)

	fmt.Fscanln(reader, &M)
	var box []int = make([]int, M)

	for i := 0; i < M; i++ {
		if i == M-1 {
			fmt.Fscanln(reader, &box[i])
		} else {
			fmt.Fscan(reader, &box[i])
		}
	}

	sort.Ints(box)

	if crane[N-1] < box[M-1] {
		fmt.Println(-1)
	} else {
		var x, y, cnt int
		x = N-1
		y = len(box)-1
		for {
			if crane[x] > box[y] {
				box = box[:y]
				x--
			} else if crane[x] < box[y] {
				y--
			}else {
				x = N-1
				cnt++
			}

			if len(box) == 0 {
				break
			}
		}

		fmt.Println(cnt)
	}
}
