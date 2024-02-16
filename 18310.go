package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var house []int = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscan(reader, &house[i])
	}

	sort.Ints(house)

	if N % 2 == 0 {
		fmt.Println(house[(N/2)-1])
	} else {
		fmt.Println(house[N/2])
	}
}