package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

func main() {
	var N, K int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &K)
	var num []int = make([]int, N)
	var ans []int = make([]int, N-K+1)

	for i:=0;i<N;i++ {
		fmt.Fscanf(reader, "%d ", &num[i])
	}

	for i:=0;i<N-K+1;i++ {
		var ansNum int
		for j:=0;j<K;j++ {
			ansNum += num[i+j] 
		}
		ans[i] = ansNum
	}

	sort.Ints(ans)

	fmt.Println(ans[N-K])
}