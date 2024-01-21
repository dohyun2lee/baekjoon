package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var K, x, sum int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var stack []int

	fmt.Fscanln(reader, &K)

	for i:=0;i<K;i++ {
		fmt.Fscanln(reader, &x)
		if x == 0 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, x)
		}		
	}

	for i:=0;i<len(stack);i++ {
		sum += stack[i]
	}

	fmt.Fprintln(writer, sum)
}