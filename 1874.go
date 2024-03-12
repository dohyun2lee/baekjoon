package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	var N, tmp int
	var sign []string
	var num map[int]int = make(map[int]int)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	stack := list.New()

	fmt.Fscanln(reader, &N)
	var ans []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &ans[i])
	}

	for i := 0; i < N; i++ {
		if i == 0 {
			for j := 1; j <= ans[i]; j++ {
				sign = append(sign, "+")
				num[j] = 1
			}
			sign = append(sign, "-")
			num[ans[0]] = 2
		} else {
			if num[ans[i]] == 1 {
				sign = append(sign, "-")
				num[ans[i]] = 2
			} else {
				for j := 1; j <= ans[i]; j++ {
					if num[j] == 0 {
						num[j] = 1
						sign = append(sign, "+")
					}
				}
				if num[ans[i]] == 1 {
					sign = append(sign, "-")
				}
			}
		}
	}
}
