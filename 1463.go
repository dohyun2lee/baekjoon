package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for {
		cnt++
		if N % 3 == 0 {
			N = N / 3
		} else if N % 2 == 0 {
			N = N /2
		} else {
			N--
		}

		fmt.Println(cnt, N)
		
		if N == 1 {
			break
		}
	}

	fmt.Println(cnt)
}