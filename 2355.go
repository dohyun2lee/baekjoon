package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var a, b int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &a, &b)

	if a > b {
		a, b = b, a
	}

	fmt.Println((b*(b+1)/2)-(a*(a-1)/2))
}