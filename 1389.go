package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, M, a, b int
	var Bacon map[int]int = make(map[int]int)
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &M)

	for i:=0;i<M;i++ {
		fmt.Fscanln(reader, &a, &b)
		
	}
}
