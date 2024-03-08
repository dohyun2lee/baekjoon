package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, M int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &M)
	var tree []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &tree[i])
	}

	sort.Ints(tree)

	fmt.Println(chk(M, tree))
}

func chk(M int, tree []int) (ans int) {
	min := 1
	max := tree[len(tree)-1]
	var mid int

	for min <= max {
		mid = (min + max) / 2
		var cnt int
		for i := 0; i < len(tree); i++ {
			if tree[i]-mid > 0 {
				cnt += tree[i] - mid
			}
		}
		if cnt >= M {
			if ans < mid {
				ans = mid
				min = mid + 1
			}
		} else {
			max = mid - 1
		}
	}
	return
}
