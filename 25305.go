package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

func main() {
	var N, k int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &k)
	var slice []int = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscan(reader, &slice[i])
	}

	sort.Ints(slice)

	fmt.Println(slice[N-k])
}