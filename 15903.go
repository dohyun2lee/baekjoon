package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, M, ans int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &M)
	var num []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &num[i])
	}

	for i := 0; i < M; i++ {
		coal(num)
	}

	for i := 0; i < N; i++ {
		ans += num[i]
	}

	fmt.Println(ans)
}

func coal (num []int) {
	sort.Ints(num)
	a := num[0]
	b := num[1]
	num[0] += b
	num[1] += a
}
