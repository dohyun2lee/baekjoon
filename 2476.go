package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, max int
	var dice []int = make([]int, 3)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &dice[0], &dice[1], &dice[2])
		sort.Ints(dice)

		if chk(dice[0], dice[1], dice[2]) > max {
			max = chk(dice[0], dice[1], dice[2])
		}
	}

	fmt.Fprintln(writer, max)
}

func chk(a, b, c int) int {
	var sum int
	if a == b && b == c {
		sum = 10000 + (a * 1000)
	} else if a == b && b != c {
		sum = 1000 + (100 * a)
	} else if a != b && b == c {
		sum = 1000 + (100 * b)
	} else if a == c && b != c {
		sum = 1000 + (100 * c)
	} else {
		sum = c * 100
	}

	return sum
}
