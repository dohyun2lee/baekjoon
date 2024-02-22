package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var tower []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &tower[i])
	}

	for i := 0; i < N-1; i++ {
		if tower[i] <= tower[i+1] {
			cnt++
		}
	}

	fmt.Println(cnt + 1)
}
