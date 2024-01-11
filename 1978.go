package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, ans int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for i:=0;i<N;i++ {
		var a, cnt int
		fmt.Fscan(reader, &a)
		
		for j:=1;j<=a;j++ {
			if a % j == 0 {
				cnt++
			}
		}
		if cnt == 2 {
			ans++
		}
	}

	fmt.Println(ans)
}