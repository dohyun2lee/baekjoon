package main

import (
	"bufio"
	"os"
	"fmt"
)

var cnt_fib, cnt_fibo int

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &n)

	fib(n)
	fibonacci(n)

	fmt.Println(cnt_fib+1, cnt_fibo)
}

func fib(n int) int {
	if n == 1 || n ==2 {
		return 1
	} else {
		cnt_fib++
		return (fib(n-1) + fib(n-2))
	}
}

func fibonacci(n int) int {
	var f []int= make([]int, n+1)
	f[1] = 1 
	f[2] = 1

	for i:=3;i<=n;i++ {
		cnt_fibo++
		f[i] = f[i-1] + f[i-2]
	}

	return f[n]
}