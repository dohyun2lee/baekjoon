package main

import (
	"fmt"
	"sort"
)

func main() {
	var basket = [5]int{0, 1, 2, 3, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(basket[1:4])))

	fmt.Println(basket)
}