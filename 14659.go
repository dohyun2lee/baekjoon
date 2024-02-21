package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, max, cnt, tmp int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var archer []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &archer[i])
	}

	tmp = archer[0]
	archer = append(archer, -1)

	for i := 0; i < N+1; i++ {
		if archer[i] == -1 {
			if max < cnt {
				max = cnt
			}
			break
		}
		if tmp > archer[i] {
			cnt++
		} else {
			if max < cnt {
				max = cnt
			}
			tmp = archer[i]
			cnt = 0
		}
	}

	fmt.Println(max)
}
