package main

import (
	"bufio"
	"fmt"
	"os"
)

var cnt0, cnt1 int

func main() {
	var N, x int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var arr []int = make([]int, 41)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &x)
		cnt0 = 0
		cnt1 = 0
		fibo(x)
		fmt.Println(cnt0, cnt1)
	}
}

func fibo(n int) int {
	if n == 0 {
		cnt0 = 1
		cnt1 = 0
		return 0
	} else if n == 1 {
		cnt1 = 1
		cnt0 = 0
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}
