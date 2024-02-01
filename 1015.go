package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var a []float64 = make([]float64, N)
	var b []float64 = make([]float64, N)
	//var ans []int = make([]int, N)
	var ans map[float64]float64 = make(map[float64]float64, N)

	for i:=0;i<N;i++ {
		fmt.Fscan(reader, &a[i])
		b[i] = a[i]
	}

	sort.Float64s(b)
	ans[b[0]] = 0

	for i:=1;i<=N;i++ {
		if b[i-1] == b[i] {
			b[i] += 0.1
		}
	}

	// for i:=0;i<N;i++ {
	// 	fmt.Fprintf(writer, "%d ", A[a[i]])
	// }
}