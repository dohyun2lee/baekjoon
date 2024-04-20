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
		fmt.Fscanf(reader, "%d ", &N)
		var grade []float64 = make([]float64, N)
		var sum, avg, num, ans float64

		for j := 0; j < N; j++ {
			if j != N-1 {
				fmt.Fscan(reader, &grade[j])
			} else {
				fmt.Fscanln(reader, &grade[j])
			}
			sum += grade[j]
		}

		avg = sum / float64(N)

		for j := 0; j < N; j++ {
			if grade[j] > avg {
				num++
			}
		}

		ans = num / float64(N) * 100
		fmt.Fprintf(writer, "%.3f", ans)
		fmt.Fprintln(writer, "%")
	}
}
