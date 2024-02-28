package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, M, price, max int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &M)
	var egg []int = make([]int, M)

	for i := 0; i < M; i++ {
		fmt.Fscanln(reader, &egg[i])
	}

	sort.Ints(egg)

	if N < M {
		egg = egg[M-N:]
	}

	for i:=0;i<len(egg);i++ {
		if max < egg[i] * (len(egg)-i) {
			max = egg[i] * (len(egg)-i)
			price = egg[i]
		}
	}

	fmt.Println(price, max)
}
