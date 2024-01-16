package main

import (
	"fmt"
	"sort"
)

func main() {
	var slice []int
	var a, sum int

	for i:=0;i<5;i++ {
		fmt.Scanln(&a)
		slice = append(slice, a)
		sum += a
	}

	sort.Ints(slice)

	fmt.Println(sum / 5)
	fmt.Println(slice[2])
}