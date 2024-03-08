package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, H, W, N, h, w int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &H, &W, &N)

		h = N % H
		w = N/H + 1

		if h == 0 {
			w--
			h = H
		}

		fmt.Fprintln(writer, h*100+w)
	}
}
