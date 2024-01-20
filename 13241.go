package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var A, B, ans int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &A, &B)

	if A > B {
		for j := 1; ; j++ {
			if (A*j)%B == 0 {
				ans = A * j
				break
			}
		}
	} else {
		for j := 1; ; j++ {
			if (B*j)%A == 0 {
				ans = B * j
				break
			}
		}
	}

	if A > 1000 && B > 1000 {
		fmt.Fprintln(writer, int64(ans))
	} else {
		fmt.Fprintln(writer, ans)
	}
}
