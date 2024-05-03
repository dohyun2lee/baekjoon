package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, cnt, x int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	N = N%10

	for i:=0;i<5;i++ {
		if i != 4 {
			fmt.Fscan(reader, &x)
		} else {
			fmt.Fscanln(reader, &x)
		}
		if x == N {
			cnt++
		}
	}

	fmt.Println(cnt)
}