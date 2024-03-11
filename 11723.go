package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var S map[int]bool = make(map[int]bool)
	var M, x int
	var cmd string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &M)

	for i := 0; i < M; i++ {
		fmt.Fscanln(reader, &cmd, &x)

		switch cmd {
		case "add":
			S[x] = true
		case "remove":
			S[x] = false
		case "check":
			if S[x] {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case "toggle":
			if S[x] {
				S[x] = false
			} else {
				S[x] = true
			}
		case "all":
			for i := 1; i <= 20; i++ {
				S[i] = true
			}
		case "empty":
			for i := 1; i <= 20; i++ {
				S[i] = false
			}
		}
	}
}
