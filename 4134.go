package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &n)
		if n < 2 {
			fmt.Fprintln(writer, "2")
		} else {
			for a := n; ; a++ {
				if primeNum(a) == true {
					fmt.Fprintln(writer, a)
					break
				}
			}
		}
	}
}

func primeNum(x int) bool {
	for i:=2;i*i<=x;i++ {
		if x%i==0 {
			return false
		}
	}
	return true
}