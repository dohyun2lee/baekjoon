package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, max, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var tree []int = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscan(reader, &tree[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(tree)))

	cnt = 2

	for i:=0;i<N;i++ {
		if cnt + tree[i] > max {
			max = cnt + tree[i]
		}
		cnt++
	}

	fmt.Println(max)
}
