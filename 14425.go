package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, M, cnt int
	var a string
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Fscanln(reader, &N, &M)
	var S = map[string]int{}

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &a)
		S[a]++
	}

	for i:=0;i<M;i++ {
		fmt.Fscanln(reader, &a)
		if S[a] != 0 {
			cnt++
		}		
	}

	fmt.Println(cnt)
}