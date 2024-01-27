package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	var s string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	s = "-"

	if N == 0 {
		fmt.Fprintln(writer, s)
	} else {
		for i:=1;i<=N;i++ {
			s = s + strings.Repeat(" ", len(s)) + s
		}
		fmt.Fprintln(writer, s)
	}
}