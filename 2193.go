package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var num []int = make([]int, 91)

	num[1] = 1
	num[2] = 1
	num[3] = 2

	if N<4 {
		fmt.Println(num[N])
	} else {
		for i:=4;i<=N;i++ {
			num[i] = num[i-2] + num[i-1]
		}

		fmt.Println(num[N])
	}
}