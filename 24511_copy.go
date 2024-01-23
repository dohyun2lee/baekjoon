package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var N, M, qs int
	var queue []int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var chk []int = make([]int, N)
	
	for i:=0;i<N;i++ {
		if i < N-1 {
			fmt.Fscanf(reader, "%d ", &chk[i])
		} else {
			fmt.Fscanln(reader, &chk[i])
		}
	}

	for i:=0;i<N;i++ {
		if i < N-1 {
			fmt.Fscanf(reader, "%d ", &qs)
		} else {
			fmt.Fscanln(reader, &qs)
		}
		if chk[i] == 0 {
			queue = append(queue, qs)
		} 
	}

	fmt.Fscanln(reader, &M)

	if chk[len(chk)-1] == 1 {
		for i:=0;i<M;i++ {
			fmt.Fscanf(reader, "%d ", &qs)
			fmt.Fprintf(writer, "%d ", qs)
		}
	} else {
		for i:=0;i<M;i++ {
			fmt.Fscanf(reader, "%d ", &qs)
			queue = append([]int{qs}, queue...)
		}

		fmt.Println("queue :", queue)

		for i:=len(queue)-1;i>=len(queue)-M;i-- {
			fmt.Fprintf(writer, "%d ", queue[i])
		}
	}
}