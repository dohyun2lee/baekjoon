package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	var cmd, x int
	var stack []int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &cmd, &x)

		switch cmd {
		case 1:
			stack = append([]int{x}, stack...)
		case 2:
			stack = append(stack, x)
		case 3:
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else if len(stack) == 1 {
				fmt.Fprintln(writer, stack[0])
				stack = []int{}
			} else {
				fmt.Fprintln(writer, stack[0])
				stack = stack[1:]
			}
		case 4:
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else if len(stack) == 1 {
				fmt.Fprintln(writer, stack[0])
				stack = []int{}
			} else {
				fmt.Fprintln(writer, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		case 5:
			fmt.Fprintln(writer, len(stack))
		case 6:
			if len(stack) == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case 7:
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, stack[0])
			}
		case 8:
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, stack[len(stack)-1])
			}
		}
	}
}
