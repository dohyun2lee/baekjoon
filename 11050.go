package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, K int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &K)
	var fact_N, fact_K, fact_nk int = 1, 1, 1

	for i := N; i > 0; i-- {
		fact_N *= i
	}
	for i := K; i > 0; i-- {
		fact_K *= i
	}
	for i := N - K; i > 0; i-- {
		fact_nk *= i
	}

	fmt.Println(fact_N / (fact_K * fact_nk))
}
