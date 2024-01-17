package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, M, a int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var card = map[int]int{}

	for i:=0;i<N;i++ {
		if i == N-1 {
			fmt.Fscanln(reader, &a)
		} else {
			fmt.Fscan(reader, &a)
		}
		card[a]++
	}

	fmt.Fscanln(reader, &M)

	for i:=0;i<M;i++ {
		fmt.Fscan(reader, &a)
		fmt.Fprintf(writer, "%d ", have(card, a))
	}
}

func have(card map[int]int, n int) int {
	if card[n] != 0 {
		return 1
	} else {
		return 0
	}
}