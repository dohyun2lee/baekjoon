package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, max, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var archer []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &archer[i])
	}

	for i := 0; i < N-1; i++ {
		cnt = 0
		for j := i+1; j < N; j++ {
			if archer[i] > archer[j] {
				cnt++
			}
		}
		if max < cnt {
			max = cnt
		}
	}

	fmt.Println(max)
}
