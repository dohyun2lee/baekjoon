package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var slice []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &slice[i])
	}

	sort.Ints(slice)

	for i := 0; i < N; i++ {
		fmt.Fprintln(writer, slice[i])
	}
}
