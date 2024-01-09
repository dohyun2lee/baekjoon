package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var T, C int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &T)

	for i:=0;i<T;i++ {
		fmt.Fscanln(reader, &C)
		var ans []int
		ans = append(ans, C / 25)
		ans = append(ans, C % 25 / 10)
		ans = append(ans, C % 25 % 10 / 5)
		ans = append(ans, C % 25 % 10 % 5)
		for j:=0;j<4;j++ {
			fmt.Printf("%d ", ans[j])
		}
	}
}