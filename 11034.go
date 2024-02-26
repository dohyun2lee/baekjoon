package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var kangaroo []int = make([]int, 3)
	reader := bufio.NewReader(os.Stdin)

	for {
		cnt, _ := fmt.Fscanf(reader, "%d %d %d\n", &kangaroo[0], &kangaroo[1], &kangaroo[2])
		if cnt != 3 {
			return
		}

		if (kangaroo[2] - kangaroo[1]) > (kangaroo[1] - kangaroo[0]) {
			fmt.Println(kangaroo[2] - kangaroo[1] - 1)
		} else {
			fmt.Println(kangaroo[1] - kangaroo[0] - 1)
		}
	}
}
