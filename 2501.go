package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var N, K, k int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &K)

	for i:=1;i<=N;i++{
		if N%i == 0 {
			k++
			if k == K {
				fmt.Println(i)
				break
			}
		}
	}
	if k < K {
		fmt.Println("0")
	}
}