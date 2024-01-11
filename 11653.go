package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)	 
	i := 2

	fmt.Fscanln(reader, &N)

	for {
		if N == 1 {
			break
		} else if N % i == 0 {
			N /= i
			fmt.Println(i)
		} else {
			i++
		}
	}
}