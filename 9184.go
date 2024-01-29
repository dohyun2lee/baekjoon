package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		fmt.Fscanln(reader, &a, &b, &c)

		if a == -1 && b == -1 && c == -1 {
			break
		}
		fmt.Fprintf(writer, "w(%d, %d, %d) = %d\n", a, b, c, w(a, b, c))
	}
}

func w(a, b, c int) int {
	if a <= 0 || b <= 0 || c <= 0 {
		return 1
	} else if a > 20 || b > 20 || c > 20 {
		return w(20, 20, 20)
	} else if a < b && b < c {
		return (w(a, b, c-1) + w(a, b-1, c-1) - w(a, b-1, c))
	} else {
		return (w(a-1, b, c) + w(a-1, b-1, c) + w(a-1, b, c-1) - w(a-1, b-1, c-1))
	}
}
