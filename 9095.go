package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, N int
	var ott []int = make([]int, 11)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	ott[0] = 0
	ott[1] = 1
	ott[2] = 2
	ott[3] = 4

	for i:=4;i<11;i++ {
		ott[i] = ott[i-3] + ott[i-2] + ott[i-1]
	}

	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &N)

		fmt.Fprintln(writer, ott[N])
	}
}
