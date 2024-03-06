package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var K, N, max int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &K, &N)
	var lan []int = make([]int, K)

	for i := 0; i < K; i++ {
		fmt.Fscanln(reader, &lan[i])

		if lan[i] > max {
			max = lan[i]
		}
	}

	fmt.Fprintln(writer, chk(N, max, lan))
}

func chk(N, max int, lan []int) (ans int) {
	min := 1
	var mid int

	for min <= max {
		mid = (min + max) / 2
		var cnt int
		for i := 0; i < len(lan); i++ {
			cnt += lan[i] / mid
		}
		if cnt >= N {
			if ans < mid {
				ans = mid
				min = mid + 1
			}
		} else {
			max = mid - 1
		}
	}
	return
}
