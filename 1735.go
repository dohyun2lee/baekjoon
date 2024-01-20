package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, A, B int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &a, &A)
	fmt.Fscanln(reader, &b, &B)

	res := (A*b + B*a)

	x := getGCD(A*B, res)

	fmt.Fprintln(writer, res/x, A*B/x)

}

func getGCD(first, second int) (gcd int) {
	if first < second {
		second, first = first, second
	}

	for second != 0 {
		first, second = second, first%second
	}
	return first
}
