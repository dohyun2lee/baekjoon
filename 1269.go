package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var N, M, cnt, a int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)
	var union map[int]int = make(map[int]int)

	for i:=0;i<N;i++ {
		if i == N-1 {
			fmt.Fscanln(reader, &a)
		} else {
			fmt.Fscan(reader, &a)
		}
		union[a]++
	}
	for i:=0;i<M;i++ {
		if i == M-1 {
			fmt.Fscanln(reader, &a)
		} else {
			fmt.Fscan(reader, &a)
		}
		union[a]++
	}
	
	for _, val := range union {
		if val == 1 {
			cnt++
		}
	}

	fmt.Fprintln(writer, cnt)
}