package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	var stack []string
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		cmd := scanner.Text()
		cmd1 := strings.Split(cmd, " ")

		if cmd1[0] == "1" {
			stack = append(stack, cmd1[1])
		} else if cmd1[0] == "2" {
			if len(stack) == 0 {
				fmt.Fprintln(writer, "-1")
			} else {
				fmt.Fprintln(writer, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		} else if cmd1[0] == "3" {
			fmt.Fprintln(writer, len(stack))
		} else if cmd1[0] == "4" {
			if len(stack) == 0 {
				fmt.Fprintln(writer, "1")
			} else {
				fmt.Fprintln(writer, "0")
			}
		} else {
			if len(stack) == 0 {
				fmt.Fprintln(writer, "-1")
			} else {
				fmt.Fprintln(writer, stack[len(stack)-1])
			}
		}
	}
}
