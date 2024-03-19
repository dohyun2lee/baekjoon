package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var a, b, c int
	var A []int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &a, &b, &c)
	A = append(A, a)
	A = append(A, b)
	A = append(A, c)

	sort.Ints(A)

	fmt.Println(A[1])
}
