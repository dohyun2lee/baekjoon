package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var M, N, ans int
	reader := bufio.NewReader(os.Stdin)
	var decimal []int

	fmt.Fscanln(reader, &M)
	fmt.Fscanln(reader, &N)

	for i:=M;i<=N;i++ {
		var cnt int
		
		for j:=1;j<=i;j++ {
			if i % j == 0 {
				cnt++
			}
		}
		if cnt == 2 {
			decimal = append(decimal, i)
			ans += i
		}
	}
	
	if ans == 0 {
		fmt.Println("-1")
	} else {
		fmt.Println(ans)
		fmt.Println(decimal[0])
	}
}