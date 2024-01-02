package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

func main(){
	var N int
	var sum = 0
	reader := bufio.NewReader(os.Stdin)

	fmt.Scanln(&N)
	var slice []int = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscan(reader, &slice[i])
		sum = sum + slice[i]
	}

	sort.Ints(slice)

	fmt.Println(float64(sum) / float64(slice[N-1]) / float64(N) * float64(100))
}