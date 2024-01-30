package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, a, b int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	
	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &a, &b)
		
		ans := 1

		for j:=0;j<b;j++ {
			ans = (ans * a) % 10
		}

		if ans == 0 {
			fmt.Println(10)
		} else {
			fmt.Println(ans)
		}
	}
}