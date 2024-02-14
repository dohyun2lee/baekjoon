package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, milk, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var store []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &store[i])
		if store[i] == milk {
			cnt++
			milk++
			if milk == 3 {
				milk = 0
			}
		}
	}

	fmt.Println(cnt)
}
