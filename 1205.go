package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, P, New, rank int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &New, &P)
	var score []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &score[i])
	}

	rank = 1
	sort.Ints(score)

	fmt.Println(score, rank)

	// for i := 0; i < N; i++ {

	// }
}
