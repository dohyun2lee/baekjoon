	package main

	import (
		"bufio"
		"fmt"
		"os"
		"sort"
	)

	func main() {
		var N int
		reader := bufio.NewReader(os.Stdin)

		fmt.Fscanln(reader, &N)
		var slice []int = make([]int, N)

		for i:=0;i<N;i++ {
			fmt.Fscan(reader, &slice[i])
		}

		sort.Ints(slice)

		fmt.Printf("%d %d", slice[0], slice[N-1])
	}