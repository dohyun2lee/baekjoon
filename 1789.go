package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, S int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &S)

	i := 1

	for S > 0 {
		if S-i >= 0 {
			N++
			S -= i
		} else {
            break
        }
        i++
	}

	fmt.Println(N)
}
