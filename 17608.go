package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var stick []int = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &stick[i])
	}

	standard := stick[N-1]

	for i:=N-1;i>=0;i-- {
		if stick[i] > standard {
			cnt++
			standard = stick[i]
		}
	}

	fmt.Println(cnt+1)
}