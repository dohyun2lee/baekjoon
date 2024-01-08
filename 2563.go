package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var matrix [100][100]int
	var N, x, y, ans int
	reader := bufio.NewReader(os.Stdin)

	fmt.Scanln(&N)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &x, &y)
		for j:=x;j<=x+10;j++ {
			for k:=y;k<=y+10;k++ {
				matrix[j][k] = 1
			}
		}
	}

	for i:=0;i<100;i++ {
		for j:=0;i<100;j++ {
			if matrix[i][j] == 1 {
				ans++
			}
		}
	}

	fmt.Println(ans)
}