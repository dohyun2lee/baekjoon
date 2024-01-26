package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

func main() {
	var N, ans, timechk int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var time []int = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscanf(reader, "%d ", &time[i])
	}

	sort.Ints(time)

	for i:=0;i<N;i++ {
		timechk += time[i]
		ans += timechk
	}

	fmt.Println(ans)
}