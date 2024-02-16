package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, ans int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var score []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &score[i])
	}

	for i := N - 1; i > 0; i-- {
		if score[i] <= score[i-1] {
			ans = ans + (score[i-1] - (score[i] - 1))
			score[i-1] = score[i] - 1
		} else {
			continue
		}
	}

	fmt.Println(ans)
}
