package main

import (
	"fmt"
	"sort"
)

func main(){
	var a = []int{1, 2, 3, 5, 4, 8, 7, 6}

	sort.Ints(a)
	fmt.Println(a)
}
