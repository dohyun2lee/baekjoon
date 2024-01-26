package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

type word struct {
	spell string
	length int
}

func main() {
	var N, M int
	var s string
	//var words map[string]int = make(map[string]int)
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &M)
	var Word []word

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &s)
		if len(s) >= M {
			Word = append(Word, word{s, len(s)})
		}
	}

	var final map[word]int = make(map[word]int)
	fmt.Println(Word)

	for i:=0;i<len(Word);i++ {
		final[Word[i]]++
	}

	sorted := make([]string, 0, len(final))

	for k:=range final {
		sorted = append(sorted, k.spell)
	}
	sort.Strings(sorted)

	for _,k:= range final {
		fmt.Println(k, final[k])
	}
	fmt.Println(final)
}
