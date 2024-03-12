package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	var cmd string
	var N, x int
	stack := list.New()
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &cmd, &x)

		switch cmd {
		case "push":
			stack.PushBack(x)
		case "pop":
			if a := stack.Back(); a != nil {
				fmt.Fprintln(writer, a.Value)
				stack.Remove(stack.Back())
			} else {
				fmt.Fprintln(writer, -1)
			}
		case "size":
			a := stack.Len()
			fmt.Fprintln(writer, a)
		case "empty":
			if a := stack.Len(); a != 0 {
				fmt.Fprintln(writer, 0)
			} else {
				fmt.Fprintln(writer, 1)
			}
		case "top":
			if a := stack.Back(); a != nil {
				fmt.Fprintln(writer, a.Value)
			} else {
				fmt.Fprintln(writer, -1)
			}
		}
	}
}
