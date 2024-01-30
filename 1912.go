package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var N, max int
	var arr []int = make([]int, 100000)
	var sum []int = make([]int, 100000)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i:=0;i<N;i++ {
		fmt.Fscanf(reader, "%d ", &arr[i])
	}

	max = arr[0]
	sum[0] = arr[0]

	for i:=1;i<N;i++ {
		sum[i] = Max((sum[i - 1] + arr[i]), arr[i])
		max = Max(sum[i], max)
	}
	fmt.Println(writer, max)
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}