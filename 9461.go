package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, T int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i:=0;i<T;i++ {
		fmt.Fscanln(reader, &N)
		fmt.Fprintln(writer, triangle(N))
	}
}

func triangle(n int) int {
	if n == 1 || n == 2 || n == 3{
		return 1
	} else if n == 4 || n == 5 {
		return 2
	} else {
		var t []int = make([]int, n+1)
		t[1] = 1
		t[2] = 1
		t[3] = 1
		t[4] = 2
		t[5] = 2
		for i := 6; i <= n; i++ {
			t[i] = (t[i-5] + t[i-1]) 
		}
		return t[n]
	}
}