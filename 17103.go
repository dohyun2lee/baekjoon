package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	N, t int
	arr  [1000001]bool
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &N)

	for i := 2; i < len(arr); i++ {
		arr[i] = true
		for j := 2; j <= int(math.Pow(float64(i), 0.5)); j++ {
			if i%j == 0 {
				arr[i] = false
				break
			}
		}
	}

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &t)
		fmt.Fprintln(writer, Goldbach(t))
	}

	writer.Flush()
}

func Goldbach(t int) int {
	count := 0
	for i := 2; i <= t/2; i++ {
		if arr[i] && arr[t-i] {
			count++
		}
	}
	return count
}