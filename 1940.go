package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, M, x, cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	fmt.Fscanln(reader, &M)
	var ingredient map[int]bool = make(map[int]bool)

	for i:=0;i<N;i++ {
		if i != N-1 {
			fmt.Fscan(reader, &x)
			ingredient[x] = true
		} else {
			fmt.Fscanln(reader, &x)
			ingredient[x] = true
		}
	}

	for key, val := range ingredient {
		if val {
			if ingredient[M-key] && key != M-key { 
				cnt++
				ingredient[key] = false
				ingredient[M-key] = false
			}
		}
	}

	fmt.Println(cnt)
}