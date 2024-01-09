package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	cnt := 1
	a := 2
	i := 6

	fmt.Fscanln(reader, &N)

	if N == 1 {
		fmt.Println("1")
	} else {
		for a <= N {
			a += i
			i += 6
			cnt++
		}
		fmt.Println(cnt)
	}
}