package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var sum, dif int
	var dwarf []int = make([]int, 9)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < 9; i++ {
		fmt.Fscanln(reader, &dwarf[i])
		sum += dwarf[i]
	}

	sort.Ints(dwarf)
	dif = sum - 100

Loop1:
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if i != j && dwarf[i]+dwarf[j] == dif {
				dwarf[i] = 0
				dwarf[j] = 0
				break Loop1
			}
		}
	}

	for i := 0; i < 9; i++ {
		if dwarf[i] != 0 {
			fmt.Fprintln(writer, dwarf[i])
		}
	}
}
