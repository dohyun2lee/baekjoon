package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var X, cnt int
	var stick = []int{64,32,16,8,4,2,1}
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &X)

	for i:=0;i<7;i++{
		if X >= stick[i] {
			cnt++
			X = X-stick[i]
			if X == 0 {
				break
			}
		}
	}

	fmt.Println(cnt)
}