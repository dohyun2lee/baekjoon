package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var a []int = make([]int, N)
	var b []int = make([]int, N)

	for i:=0;i<N;i++ {
		if i != N-1 {
			fmt.Fscan(reader, &a[i])
		} else {
			fmt.Fscanln(reader, &a[i])
		}
	}

	for i:=0;i<N;i++ {
		if i != N-1 {
			fmt.Fscan(reader, &b[i])
			if b[i]-a[i] > 0 {
				cnt += b[i] - a[i]
			}
		} else {
			fmt.Fscanln(reader, &b[i])
			if b[i]-a[i] > 0 {
				cnt += b[i] - a[i]
			}
		}
	}

	fmt.Println(cnt)
}
