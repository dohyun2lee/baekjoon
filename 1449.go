package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, L, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &L)
	var water []int = make([]int, N)
	var waterProof []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &water[i])
		waterProof[i] = -1
	}

	sort.Ints(water)

	for i := 0; i < N; i++ {
		if waterProof[i] == -1 {
			cnt++
		} else {
			continue
		}
		for j:=i;j<N;j++ {
			if water[i] + L > water[j] {
				waterProof[j] *= -1
			} else {
				break
			}
		}
	}

	fmt.Println(cnt)
}
