package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, M, x int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for ii := 0; ii < T; ii++ {
		fmt.Fscanln(reader, &M)
		var num map[int]bool = make(map[int]bool, M)

		for i := 0; i < M; i++ {
			if i != M-1 {
				fmt.Fscan(reader, &x)
			} else {
				fmt.Fscanln(reader, &x)
			}

			num[x] = true
		}

		for i := 0; i < M; i++ {
			if i != M-1 {
				fmt.Fscan(reader, &x)
			} else {
				fmt.Fscanln(reader, &x)
			}

			if num[x] {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		}
	}
}
