package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var rank []int = make([]int, N+1)

	for i := 1; i <= N; i++ {
		fmt.Fscanln(reader, &rank[i])
	}

	sort.Ints(rank)

	for i:=1;i<=N;i++ {
		cnt += (i - rank[i])
	}

	fmt.Println(cnt)
}
