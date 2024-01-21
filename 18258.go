package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	var N int
	var stack []int
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		cmd := scanner.Text()
		cmd1 := strings.Split(cmd, " ")

		switch cmd1[0] {
		case "push" :
			x, _ := strconv.Atoi(cmd1[1])
			stack = append(stack, x)
		case "pop" :
			if len(stack) == 0 {
				fmt.Fprintln(writer, -1)
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
