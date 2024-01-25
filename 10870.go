package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	if N == 0 {
		fmt.Fprintln(writer, 0)
	} else {
		var arr []int = make([]int, N+1)

		arr[0] = 0
		arr[1] = 1
	
		for i:=2;i<=N;i++ {
			arr[i] = arr[i-2] + arr[i-1]
		}

		fmt.Fprintln(writer, arr[N])
	}
}