package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &t)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0 ; i < t; i++ {
		var n int
		fmt.Fscanln(reader, &n)
		cntzero, cntone := result(n)
		fmt.Fprintf(writer, "%d %d\n", cntzero[n], cntone[n])
	}
}

func result(n int) (cntzero, cntone[]int) {
	cntzero = []int{1, 0}
	cntone = []int{0, 1}

	if n <= 1 {
		return
	}
	for i := 2; i< n+1; i++ {
		cntzero = append(cntzero, cntzero[i-1]+cntzero[i-2])
		cntone = append(cntone, cntone[i-1]+cntone[i-2])
	}
	return
}