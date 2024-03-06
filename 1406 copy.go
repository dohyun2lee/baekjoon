package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var S, x, cmd string
	var M, cursor int
	var tmp []string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &S)
	fmt.Fscanln(reader, &M)

	s := strings.Split(S, "")
	cursor = len(s) - 1

	for i := 0; i < M; i++ {
		fmt.Fscanln(reader, &cmd, &x)

		switch cmd {
		case "L":
			if cursor != -1 {
				cursor--
			}
		case "D":
			if cursor != len(s)-1 {
				cursor++
			}
		case "B":
			if cursor != -1 {
			tmp = []string{}
			tmp = append(tmp, s[:cursor]...)
			tmp = append(tmp, s[cursor+1:]...)
			s = tmp
			cursor--
			}
		case "P":
			tmp = []string{}
			tmp = append(tmp, s[:cursor+1]...)
			tmp = append(tmp, x)
			tmp = append(tmp, s[cursor+1:]...)
			s = tmp
			cursor++
		}
	}

	S = strings.Join(s, "")

	fmt.Fprintln(writer, S)
}
