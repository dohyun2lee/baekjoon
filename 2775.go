package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var person [15][15]int
	var T, K, N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i:=1;i<15;i++ {
		person[0][i] = i
	}
	for i:=1;i<15;i++{
		for j:=1;j<15;j++{
			person[i][j] = person[i-1][j]+person[i][j-1]
		}
	}

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &K)
		fmt.Fscanln(reader, &N)

		fmt.Fprintln(writer, person[K][N])
	}

}
