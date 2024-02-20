package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var a []int = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscan(reader, &a[i])
	}
	
	fmt.Println(a)
}
