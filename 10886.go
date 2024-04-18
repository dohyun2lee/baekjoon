package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, vote, res int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &vote)

		if vote == 0 {
			res--
		} else {
			res++
		}
	}

	if res > 0 {
		fmt.Println("Junhee is cute!")
	} else {
		fmt.Println("Junhee is not cute!")
	}
}
