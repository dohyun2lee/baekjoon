package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, len int
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
			len++
		case 2:
			stack = append(stack, x)
			len++
		case 3:
			if len == 0 {
				fmt.Fprintln(writer, -1)
			} else if len == 1 {
				fmt.Fprintln(writer, stack[0])
				stack = []int{}
				len--
			} else {
				fmt.Fprintln(writer, stack[0])
				stack = stack[1:]
				len--
			}
		case 4:
			if len == 0 {
				fmt.Fprintln(writer, -1)
			} else if len == 1 {
				fmt.Fprintln(writer, stack[0])
				stack = []int{}
				len--
			} else {
				fmt.Fprintln(writer, stack[len-1])
				stack = stack[:len-1]
				len--
			}
		case 5:
			fmt.Fprintln(writer, len)
		case 6:
			if len == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case 7:
			if len == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, stack[0])
			}
		case 8:
			if len == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, stack[len-1])
			}
		}
	}
}
