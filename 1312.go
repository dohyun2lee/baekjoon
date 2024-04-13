package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var a, b, n int
	res := 0
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &a, &b, &n)

	a = a % b

	for i := 0; i < n; i++ {
		a *= 10
		res = int(a / b)
		a = a % b
	}

	fmt.Printf("%d\n", res)
}