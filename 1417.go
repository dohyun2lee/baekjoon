package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, cnt, x, vote, ft int
	var candidate map[int]int = make(map[int]int)
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &vote)
		candidate[i] = vote
	}


	if N == 1 {
		cnt = 0
	} else {
		x = -1
		for {
			x = find(candidate)
			if 0 == x {
				break
			}
			candidate[0] += 1
			candidate[x] -= 1
			cnt++
		}

		y := candidate[0]
		candidate[0] = -1

		for _, val := range candidate {
			if val == y {
				ft = 1
			}
		}
	}

	fmt.Println(cnt + ft)
}

func find(candidate map[int]int) int {
	var max, cand int
	for key, val := range candidate {
		if max < val {
			max = val
			cand = key
		}
	}
	return cand
}
