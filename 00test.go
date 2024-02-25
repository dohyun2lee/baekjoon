package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type water struct {
	start, end int
}

func main() {
	var N, L, a, b int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &L)
	var pool []water

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &a, &b)
		pool = append(pool, water{a, b})
	}

	sort.Slice(pool, func(i, j int) bool {
		return pool[i].start < pool[j].start
	})

	pool[0].start = 0

	fmt.Println(pool)
}
