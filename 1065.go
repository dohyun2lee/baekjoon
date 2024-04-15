package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, cnt, hun, ten, one int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	if N < 100 {
		fmt.Println(N)
	} else {
		cnt = 99
		for i := 100; i <= N; i++ {
			hun = i / 100      
			ten = (i / 10) % 10 
			one = i % 10

			if (hun - ten) == (ten - one) { 
				cnt++
			}
		}

		fmt.Println(cnt)
	}
}
