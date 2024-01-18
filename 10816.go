package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var N, M int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var cards []int = make([]int, N)
	var sheet map[int]int = make(map[int]int)

	for i:=0;i<N;i++ {
		if i == N -1 {
			fmt.Fscanln(reader, &cards[i])
		} else {
			fmt.Fscan(reader, &cards[i])
		}
		sheet[cards[i]]++
	}

	fmt.Fscanln(reader, &M)
	var que []int = make([]int, M)
	
	for i:=0;i<M;i++ {
		fmt.Fscanf(reader, "%d ", &que[i])
	}

	for i:=0;i<M;i++ {
		fmt.Fprintf(writer, "%d ", sheet[que[i]])
	}
}