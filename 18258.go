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
		var cmd string
		var x int
		fmt.Fscanln(reader, &cmd, &x)

		switch cmd {
		case "push" :
			stack = append(stack, x)
		case "pop" :
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else if len(stack) == 1 {
				fmt.Fprintln(writer, stack[0])
				stack = []int{}
			} else {
				fmt.Fprintln(writer, stack[0])
				stack = stack[1:]
			}
		case "size" :
			fmt.Fprintln(writer, len(stack))
		case "empty" :
			if len(stack) == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case "front" :
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, stack[0])
			}
		case "back" :
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, stack[len(stack)-1])
			}
		}
	}
}
