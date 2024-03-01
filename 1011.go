package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, x, y int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		var cnt, length, move, sum int

		fmt.Fscanln(reader, &x, &y)

		length = y - x
		move = 1

		if length == 1 {
			cnt = 1
		} else if length == 2 {
			cnt = 2
		} else {
			for {
				if length <= sum {
					break
				}
				sum += move
				cnt++
				if length <= sum {
					break
				}
				sum += move
				cnt++
				if length <= sum {
					break
				}
				move++
			}
		}

		fmt.Fprintln(writer, cnt)
	}
}
