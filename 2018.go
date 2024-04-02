package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, l, r, sum, ans int = 0, 1, 1, 0, 0
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for l <= r && r <= N {

	}
}
