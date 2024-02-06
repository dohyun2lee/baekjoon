package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, x int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < 3; i++ {
		var num []int
		var sum int

		fmt.Fscanln(reader, &N)

		for j := 0; j < N; j++ {
			fmt.Fscanln(reader, &x)
			num = append(num, x)
		}

		sort.Ints(num)

		for k:=0;k<len(num)/2;k++ {
			sum += num[k] + num[len(num)-k-1]
		}
		
		if len(num)%2 != 0 {
			sum += num[len(num)/2]
		}

		if sum == 0 {
			fmt.Fprintln(writer, "0")
		} else if sum > 0 {
			fmt.Fprintln(writer, "+")
		} else {
			fmt.Fprintln(writer, "-")
		}
	}
}
