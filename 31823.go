package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var S, name string
	var N, M, max, cnt int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)

	for i:=0;i<N;i++ {
		max = 0 
		cnt = 0
		fmt.Fscanln(reader, &S, &name)

		for j:=0;j<M;j++ {
			if S[j] != 42 {
				cnt++
			} else {
				if cnt > max {
					max = cnt
				}
				cnt = 0
			}
		}

		fmt.Fprintln(writer, max, name)
	}
}

