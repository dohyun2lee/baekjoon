package main  

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	var slice []int

	fmt.Scanln(&N)

	for i:=0;i<N;i++ {
		var a int
		fmt.Scanln(&a)
		slice = append(slice, a)
	}

	sort.Ints(slice)

	for i:=0;i<N;i++ {
		fmt.Println(slice[i])
	}
}