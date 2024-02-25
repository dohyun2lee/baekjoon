package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, ans int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var line []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &line[i])
	}

	sort.Ints(line)

	ans = -1

	for i := N - 3; i >= 0; i-- {
		if (line[i] + line[i+1]) > line[i+2] {
			ans = line[i] + line[i+1] + line[i+2]
			break
		}
	}

	fmt.Println(ans)
}
