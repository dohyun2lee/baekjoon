package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var boy []int = make([]int, 6)
	var girl []int = make([]int, 6)
	var cnt, N, K, S, Y int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &K)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &S, &Y)
		if S == 1 {
			boy[Y-1]++
		} else {
			girl[Y-1]++
		}
	}

	for i:=0;i<6;i++ {
		if boy[i]%K != 0 {
			cnt += boy[i]/K + 1
		} else {
			cnt += boy[i]/K
		}
		if girl[i]%K != 0 {
			cnt += girl[i]/K + 1
		} else {
			cnt += girl[i]/K
		}
	}

	fmt.Println(cnt)
}
