package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var N, x int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	N *= 2

	x = int(math.Sqrt(float64(N)))

	fmt.Println(4 * x)
}
