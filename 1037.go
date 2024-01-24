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
	var divisor []int = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscanf(reader, "%d ", &divisor[i])
	}

	sort.Ints(divisor)

	fmt.Println(divisor[0] * divisor[len(divisor)-1])
}

