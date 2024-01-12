package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

func main() {
	var N, X, Y int
	var x []int
	var y []int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &X, &Y)
		x = append(x, X)
		y = append(y, Y)
	}

	sort.Ints(x) 
	sort.Ints(y)
	
	fmt.Println((x[N-1] - x[0]) * (y[N-1]- y[0]))
}