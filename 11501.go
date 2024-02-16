package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &N)
		var money, max int
		var price []int = make([]int, N)

		for i := 0; i < N; i++ {
			if i == N-1 {
				fmt.Fscanln(reader, &price[i])
			} else {
				fmt.Fscan(reader, &price[i])
			}
		}

		for i := N-1; i >= 0; i-- {
			if max < price[i] {
				max = price[i]
			} else {
				money += max - price[i]
			}
		}

		fmt.Fprintln(writer, money)
	}
}
