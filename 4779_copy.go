package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input int
	var s string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		_, err := fmt.Fscanln(reader, &input)
		if err == nil {
			s = "-"

			if input == 0 {
				fmt.Fprintln(writer, s)
			} else {
				for i := 1; i <= input; i++ {
					s = s + strings.Repeat(" ", len(s)) + s
				}
				fmt.Fprintln(writer, s)
			}
		} else {
			break
		}
	}
}
