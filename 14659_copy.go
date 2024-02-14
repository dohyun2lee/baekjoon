package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, max, cnt, x int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var archer []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &archer[i])
	}

	for i := 0; i < N; i++ {
		if x < archer[i] {
			x = archer[i]
			if cnt > max {
				max = cnt
				cnt = 0
			}
		} else {
			cnt++
		}	
	}

	if max < cnt {
		max = cnt
	}

	fmt.Println(max)
}
