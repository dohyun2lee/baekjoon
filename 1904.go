package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	fmt.Fprintln(writer, fibonacci(N))
}

func fibonacci(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		var f []int = make([]int, n+1)
		f[1] = 1
		f[2] = 2
		for i := 3; i <= n; i++ {
			f[i] = (f[i-1] + f[i-2]) % 15746
		}
		return f[n]
	}
}
