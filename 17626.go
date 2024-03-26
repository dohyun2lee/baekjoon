package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var N, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	n := int(math.Sqrt(float64(N)))

	for {
		if N >= (n * n) {
			fmt.Println(N)
			cnt++
			N -= n * n
		} else {
			n--
		}

		if N == 0 {
			break
		}
	}

	fmt.Println(cnt)
}
