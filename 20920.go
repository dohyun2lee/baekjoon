package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

type word struct {
	spell string
	cnt int
}
func main() {
	var N, M int
	var s string
	var tmp map[string]int = make(map[string]int)
	var input []word
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)

	for i:=0;i<N;i++ {
		fmt.Fscan(reader, &s)
		if len(s) >= M {
			tmp[s]++
		}
	}

	for key, val := range tmp {
		input = append(input, word{key, val})
	}

	sort.Slice(input, func(i, j int) bool {
		if input[i].cnt == input[j].cnt {
			if len(input[i].spell) == len(input[j].spell) {
				return input[i].spell < input[j].spell
			}
			return len(input[i].spell) > len(input[j].spell)
		}
		return input[i].cnt > input[j].cnt
	})

	for _, val := range input {
		fmt.Fprintln(writer, val.spell)
	}
}