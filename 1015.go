package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var a []int = make([]int, N)
	var b []int = make([]int, N)
	//var ans []int = make([]int, N)
	var A map[int]int = make(map[int]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscan(reader, &a[i])
		b[i] = a[i]
	}

	sort.Ints(b)

	for i:=1;i<=N;i++ {
		if A[b[i]] > 0
		A[b[i]] = i
	}

	for i:=0;i<N;i++ {
		fmt.Fprintf(writer, "%d ", A[a[i]])
	}
}