package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		var sum int
		fmt.Fscanln(reader, &N)

		if N == "0" {
			break
		}
		n := strings.Split(N, "")

		for i := 0; i < len(n); i++ {
			switch n[i] {
			case "1":
				sum += 2
			case "0":
				sum += 4
			default:
				sum += 3
			}
			sum += 1
		}

		fmt.Fprintln(writer, sum+1)
	}
}
