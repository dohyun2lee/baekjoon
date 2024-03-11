package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c, d, e int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &a, &b, &c, &d, &e)

	sum := (a * a) + (b * b) + (c * c) + (d * d) + (e * e)

	fmt.Println(sum%10)
}
