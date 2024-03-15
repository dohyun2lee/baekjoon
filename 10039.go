package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var score, sum int
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 5; i++ {
		fmt.Fscanln(reader, &score)
		if score < 40 {
			sum += 40
		} else {
			sum += score
		}
	}

	fmt.Println(sum / 5)
}
