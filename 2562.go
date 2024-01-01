package main

import (
	"sort"
	"fmt"
)

func main() {
	var one []int 
	var two []int

	for i:=0;i<9;i++ {
		a := 0
		fmt.Scanln(&a)
		one = append(one, a)
		two = append(two, a)
	}

	sort.Ints(two)

	cnt := 1

	for i:=0;i<9;i++ {
		if one[i] == two[8] {
			break
		}
		cnt++
	}

	fmt.Println(two[8], cnt)
}