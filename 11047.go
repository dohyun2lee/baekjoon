package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, K, chk, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &K)
	var coin []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &coin[i])
		if coin[i] > K {
			chk = i - 1
		}
	}

	if chk == 0 {
		chk = N - 1
	}

	for i := chk; i >= 0; i-- {
		cnt += K / coin[i]
		K = K % coin[i]
		if K == 0 {

		}
	}

	fmt.Println(cnt)
}
