package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, L int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &L)
	var froot []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &froot[i])
	}

	sort.Ints(froot)

	for i := 0; i < N; i++ {
		if L >= froot[i] {
			L++
		} else {
			break
		}
	}

	fmt.Println(L)
}
