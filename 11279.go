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
				sort.Ints(num)
				
			}
		}
	}
