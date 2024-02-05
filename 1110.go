package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, cnt, old, new int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	old = N

	for {
		new = (old%10)*10 + ((old/10)+(old%10))%10
		cnt++
		if new == N {
			break
		}
		old = new
	}

	fmt.Println(cnt)
}