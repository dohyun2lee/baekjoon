package main 

import (
	"fmt"
	"sort"
)

func main() {
	var a int
	var cnt = 10
	var remainder []int

	for i:=0;i<10;i++ {
		fmt.Scanln(&a)
		remainder = append(remainder, a%42)
	}

	sort.Ints(remainder)

	for i:=0;i<9;i++ {
		if remainder[i] == remainder[i+1] {
			cnt = cnt - 1
		}
	}

	fmt.Println(cnt)
}