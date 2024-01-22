package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	var stack []int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		var cmd, x int
		fmt.Fscanln(reader, &cmd, &x)

		switch cmd {
		case 1:
			stack = append(stack, x)
		case 2:
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else if len(stack) == 1 {
				fmt.Fprintln(writer, stack[len(stack)-1])
				stack = []int{}
			} else {
				fmt.Fprintln(writer, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		case 3:
			fmt.Fprintln(writer, len(stack))
		case 4:
			if len(stack) == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case 5:
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, stack[len(stack)-1])
			}
		}
	}
}
