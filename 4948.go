package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		cnt := 0
		fmt.Fscanln(reader, &n)

		if n == 0 {
			break
		} else if n == 1 {
			fmt.Fprintln(writer, "1")
		} else {
			for j:=n+1;j<=2*n;j++ {
				if primeNum(j) == true {
					cnt++
				}
			}
			fmt.Fprintln(writer, cnt)
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