package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var M int
	var ans float64
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &M)

	if M <= 30 {
		ans = float64(M) / 2
	} else {
		ans = float64(M-30)*3/2 + 15
	}

	fmt.Printf("%0.1f\n", ans)
}
