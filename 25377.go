package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, a, b, ans int
	reader := bufio.NewReader(os.Stdin)

	ans = 1001

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &a, &b)
		if a <= b {
			if b < ans {
				ans = b
			}
		}
	}

	if ans == 1001 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
