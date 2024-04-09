package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var max, in, out, sum int
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 4; i++ {
		fmt.Fscanln(reader, &out, &in)

		sum -= out
		sum += in

		if sum > max {
			max = sum
		}
	}

	fmt.Println(max)
}
