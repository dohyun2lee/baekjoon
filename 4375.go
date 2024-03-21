package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		cnt, _ := fmt.Fscanln(reader, &N)

		if cnt != 1 {
			break
		}
		
		var digit, num int = 1, 1

		for {
			if num % N == 0 {
				fmt.Fprintln(writer, digit)
				break
			}
			num = (num*10 + 1) % N
			digit++
		}
	}
}
