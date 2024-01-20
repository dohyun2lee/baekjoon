package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int 
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var tree []int = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &tree[i])
	}
	gcd := tree[1] - tree[0]

	for i:=1;i<N-1;i++ {
		gcd = getGCD(gcd, tree[i+1]-tree[i])
	}

	fmt.Println((tree[N-1]-tree[0])/gcd-N+1)
}

func getGCD(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a , b = b, a%b
	}
	return a
}