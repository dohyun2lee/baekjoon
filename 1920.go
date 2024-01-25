package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

func main() {
	var N, M, m int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var input []int = make([]int, N)

	for i:=0;i<N;i++ {
		if i == N-1 {
			fmt.Fscanln(reader, &input[i])
		} else {
			fmt.Fscanf(reader, "%d ", &input[i])
		}
	}
	
	sort.Ints(input)

	fmt.Fscanln(reader, &M)

	for i:=0;i<M;i++ {
		fmt.Fscanf(reader, "%d ", &m)
		fmt.Fprintln(writer, binary(input, m))
	}
}

func binary(arr []int, x int) int {
	var start, end int
	
	start = 0
	end = len(arr) - 1

	for start <= end {
		mid := (start + end) / 2
		if arr[mid] > x {
			end = mid - 1
		} else if arr[mid] < x {
			start = mid + 1
		} else {
			return 1
		}
	}
	return 0
}