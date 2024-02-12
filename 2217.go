package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, max int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var rope []int = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &rope[i])
	}

	sort.Ints(rope)

	for i:=0;i<N;i++ {
		if (rope[i] * (N - i)) > max {
			max = rope[i] * (N - i)
		}
	}
	fmt.Println(max)
}
