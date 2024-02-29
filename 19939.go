package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, K, tmp int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &K)

	tmp = (K * (K + 1)) / 2

	if N < tmp {
		fmt.Println(-1)
	} else {
		if (N-tmp)%K == 0 {
			fmt.Println(K - 1)
		} else {
			fmt.Println(K)
		}
	}
}
