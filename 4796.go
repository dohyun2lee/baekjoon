package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var L, P, V, ans, i int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		i++
		fmt.Fscanln(reader, &L, &P, &V)
		if L == 0 && P == 0 && V == 0 {
			break
		}
		tmp := (V / P) * L
		if (V % P) <= L {
			ans = tmp + (V % P)
		} else {
			ans = tmp + L
		}
		fmt.Fprintf(writer, "Case %d: %d\n", i, ans)
	}
}
