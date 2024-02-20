package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"math"
)

func main() {
	var T, N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for z := 0; z < T; z++ {
		fmt.Fscanln(reader, &N)
		var max int
		var log []int = make([]int, N)
		var ansLog []int = make([]int, N)

		for i := 0; i < N; i++ {
			if i == N-1 {
				fmt.Fscanln(reader, &log[i])
			} else {
			fmt.Fscan(reader, &log[i])
			}
		}

		sort.Ints(log)
		x := 0
		y := N - 1

		for i := 0; i < N; i++ {
			if i%2 == 0 {
				ansLog[x] = log[i]
				x++
			} else {
				ansLog[y] = log[i]
				y--
			}
		}

		ansLog = append(ansLog, ansLog[0])

		for i := 0; i < N; i++ {
			if int(math.Abs(float64(ansLog[i+1]-ansLog[i]))) > max {
				max = int(math.Abs(float64(ansLog[i+1]-ansLog[i])))
			}
		}
		fmt.Fprintln(writer, max)
	}
}
