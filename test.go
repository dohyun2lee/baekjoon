package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer
)

func init() {
	rd = bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()

	var s string
	fmt.Fscan(rd, &s)

	d := [10]int{}
	var c int
	fmt.Println("s : ", s)
	for _, v := range s {
		fmt.Println("v : ", v)
		fmt.Println("v-0 : ", v-'0')
		d[v-'0']++
		c += int(v - '0')
	}

	if c%3 > 0 || d[0] == 0 {
		fmt.Fprintf(wr, "%d", -1)
		return
	}

	fmt.Println(d)

	for i := 9; i >= 0; i-- {
		for j := 0; j < d[i]; j++ {
			fmt.Fprintf(wr, "%d", i)
		}
	}
}
