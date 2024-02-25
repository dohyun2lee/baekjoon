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
	var N, L, a, b, cnt, tmp int
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

	for i := 0; i < N; i++ {
		if tmp >= pool[i].end {
			continue
		} else {
			if tmp >= pool[i].start {
				pool[i].start = tmp
			}
			if (pool[i].end-pool[i].start)%L == 0 {
				cnt += (pool[i].end - pool[i].start) / L
				tmp = pool[i].end
			} else {
				cnt += (pool[i].end-pool[i].start)/L + 1
				tmp = pool[i].end - ((pool[i].end - pool[i].start) % L) + L
			}
		}
	}

	fmt.Println(cnt)
}
