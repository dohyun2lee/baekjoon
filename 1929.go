package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var M, N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &M, &N)

	if M == 1 {
		for a := 2; a <= N; a++ {
			if primeNum(a) == true {
				fmt.Fprintln(writer, a)
			}
		}
	} else {
		for a := M; a <= N; a++ {
			if primeNum(a) == true {
				fmt.Fprintln(writer, a)
			}
		}
	}
}

func primeNum(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
