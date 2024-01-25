package main

import (
	"bufio"
	"fmt"
	"os"
)

var cnt int

func main() {
	var T int
	var s string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		cnt = 0
		fmt.Fscanln(reader, &s)
		x := isPalindrome(s)
		fmt.Fprintln(writer, x, cnt)
	}
}

func isPalindrome(s string) int {
	x := recursion(s, 0, len(s)-1)
	return x
}

func recursion(s string, a, b int) int {
	cnt++
	if a >= b {
		return 1
	} else if s[a] != s[b] {
		return 0
	} else {
		return recursion(s, a+1, b-1)
	}
}
