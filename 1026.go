package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var S, N int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var A []int = make([]int, N)
	var b []int = make([]int, N)
	var B map[int]int = make(map[int]int, N)

	for i := 0; i < N; i++ {
		if i == N-1 {
			fmt.Fscanln(reader, &A[i])
		} else {
			fmt.Fscan(reader, &A[i])
		}
	}

	for i := 0; i < N; i++ {
		if i == N-1 {
			fmt.Fscanln(reader, &b[i])
		} else {
			fmt.Fscan(reader, &b[i])
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	sort.Ints(b)

	for i:=0;i<N;i++ {
		B[b[i]] = i
	}

	for key, val := range B {
		S += key * A[val]
	}

	fmt.Println(S)
}
