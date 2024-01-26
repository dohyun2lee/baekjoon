	package main

	import (
		"bufio"
		"fmt"
		"os"
		"sort"
	)

	func main() {
		var N, a int
		var num []int
		reader := bufio.NewReader(os.Stdin)

		fmt.Fscanln(reader, &N)

		for i := 0; i < N; i++ {
			fmt.Fscanln(reader, &a)
			if a == 0 {
				if len(num) == 0 {
					fmt.Println(0)
				} else if len(num) == 1 {
					fmt.Println(num[0])
					num = []int{}
				} else {
					sort.Ints(num)
					fmt.Println(num[len(num)-1])
					num = num[:len(num)-1]
				}
			} else {
				num = append(num, a)
			}
		}
	}
