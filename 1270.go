package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		var ans int
		fmt.Fscan(reader, &N)
		var soldier []int = make([]int, N)
		var percentage map[int]int = make(map[int]int)

		for j := 0; j < N; j++ {
			if j != N-1 {
				fmt.Fscan(reader, &soldier[j])
			} else {
				fmt.Fscanln(reader, &soldier[j])
			}
			percentage[soldier[j]]++
		}

		for key, val := range percentage {
			if val > N/2 {
				ans = key
				break
			}
		}

		if ans == 0 {
			fmt.Fprintln(writer, "SYJKGW")
		} else {
			fmt.Fprintln(writer, ans)
		}
	}
}
