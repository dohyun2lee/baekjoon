package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, sum, min1, min2, min3, sum1, sum2, sum3, a int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var dice []int
	var min []int

	for i := 0; i < 6; i++ {
		fmt.Fscan(reader, &a)
		dice = append(dice, a)
	}

	if dice[0] < dice[5] {
		min1 = dice[0]
	} else {
		min1 = dice[5]
	}

	min = append(min, min1)

	if dice[1] < dice[4] {
		min2 = dice[1]
	} else {
		min2 = dice[4]
	}

	min = append(min, min2)

	if dice[2] < dice[3] {
		min3 = dice[2]
	} else {
		min3 = dice[3]
	}

	min = append(min, min3)
	sort.Ints(min)

	if N == 1 {
		sort.Ints(dice)
		for i := 0; i < 5; i++ {
			sum += dice[i]
		}
	} else {
		sum1 = min[0]
		sum2 = min[0] + min[1]
		sum3 = min[0] + min[1] + min[2]
		sum += sum1 * ((5 * N * N) - (16 * N) + 12)
		sum += sum2 * ((8 * N) - 12)
		sum += sum3 * 4
	}

	fmt.Println(sum)
}
