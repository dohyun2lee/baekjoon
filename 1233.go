package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

type dice struct {
	num, cnt int
}

func main() {
	var a, b, c int
	var sum map[int]int = make(map[int]int)
	var Dice []dice
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &a, &b, &c)

	for i:=1;i<=a;i++ {
		for j:=1;j<=b;j++ {
			for k:=1;k<=c;k++ {
				sum[i+j+k]++
			}
		}
	}

	for key, val := range sum {
		Dice = append(Dice, dice{key, val})
	}

	sort.Slice(Dice, func(i, j int) bool {
		if Dice[i].cnt == Dice[j].cnt {
			return Dice[i].num < Dice[j].num
		}
		return Dice[i].cnt > Dice[j].cnt
	})

	fmt.Println(Dice[0].num)
}