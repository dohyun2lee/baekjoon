package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var name string
	var N, M, max, cnt int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)

	for i:=0;i<N;i++ {
		var reSt []string = make([]string, M)
		max = 0 
		cnt = 0

		for j:=0;j<M;j++ {
			fmt.Fscan(reader, &reSt[j])

			if string(reSt[j]) == "." {
				cnt++
			} else {
				if cnt > max {
					max = cnt
				}
				cnt = 0
			}
		}

		fmt.Fscanln(reader, &name)

		fmt.Fprintln(writer, max, name)
	}
}

