package main

import (
	"fmt"
	"sort"
)

func main() {
	var sum int
	var length []int = make([]int, 3)

	fmt.Scanln(&length[0], &length[1], &length[2])

	sort.Ints(length)

	if length[0] + length[1] > length[2] {
		sum = length[0] + length[1] + length[2]
	} else {
		sum = length[0] + length[1] + length[0] + length[1] - 1
	}

	fmt.Println(sum)
}