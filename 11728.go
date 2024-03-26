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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)
	var arr []int = make([]int, N+M)

	for i := 0; i < N; i++ {
		if i != N-1 {
			fmt.Fscan(reader, &arr[i])
		} else {
			fmt.Fscanln(reader, &arr[i])
		}
	}
	for i := N; i < N+M; i++ {
		if i != N+M-1 {
			fmt.Fscan(reader, &arr[i])
		} else {
			fmt.Fscanln(reader, &arr[i])
		}
	}

	sort.Ints(arr)

	for i := 0; i < N+M; i++ {
		fmt.Fprintf(writer, "%d ", arr[i])
	}
}
