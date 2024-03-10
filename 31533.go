package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, m, n, ans float64
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &m, &n)

	if m > n {
		n /= a
		if n*2 > m+n {
			ans = m + n
		} else {
			ans = n * 2
		}
	} else {
		m /= a
		if m*2 > m+n {
			ans = m + n
		} else {
			ans = m * 2
		}
	}

	fmt.Println(ans)
}
