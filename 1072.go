package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var X, Y, Z, ans, right, left, mid int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &X, &Y)
	Z = (Y * 100) / X

	if Z >= 99 {
		ans = -1
	} else {
		left = 1
		right = X
		for left <= right {
			mid = (left + right) / 2
			if (((Y + mid) * 100) / (X + mid)) > Z {
				ans = mid
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	fmt.Fprintln(writer, ans)
}
