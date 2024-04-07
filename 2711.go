package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, x int
	var word, ans string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		ans = ""
		fmt.Fscanln(reader, &x, &word)

		for j := 0; j < len(word); j++ {
			if j != x-1 {
				ans += string(word[j])
			}
		}

		fmt.Fprintln(writer, ans)
	}
}
