package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, y, m int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var time []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &time[i])
		y += (time[i]/30 + 1) * 10
		m += (time[i]/60 + 1) * 15
	}

	if y > m {
		fmt.Fprintf(writer, "M %d", m)
	} else if m > y {
		fmt.Fprintf(writer, "Y %d", y)
	} else {
		fmt.Fprintf(writer, "Y M %d", m)
	}
}
