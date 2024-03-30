package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var N, sum int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	var grade []int = make([]int, N)
	fif := int(math.Round(float64(N) * (0.15)))

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &grade[i])
	}

	sort.Ints(grade)

	grade = grade[:N-fif]
	grade = grade[fif:]

	for i := 0; i < len(grade); i++ {
		sum += grade[i]
	}

	fmt.Fprintln(writer, math.Round(float64(sum)/float64(len(grade))))
}
