package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	ans := 1

	for i := N; i > 0; i-- {
		ans *= i
	}

	fmt.Println(ans)
}
