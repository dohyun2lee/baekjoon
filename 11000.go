package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var cnt int

func main() {
	var N, a, b int
	var lecture [][]int = make([][]int, N)
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		var time []int = make([]int, 2)
		fmt.Fscanln(reader, &a, &b)
		time[0] = a
		time[1] = b
		lecture = append(lecture, time)
	}

	sort.Slice(lecture, func(i, j int) bool {
		if lecture[i][0] == lecture[j][0] {
			return lecture[i][1] < lecture[j][1]
		}
		return lecture[i][0] < lecture[j][0]
	})

	fmt.Println(lecture)
}

func chkTime(lecture [][]int) {
	
}
