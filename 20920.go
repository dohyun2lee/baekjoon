package main

import (
	"bufio"
	"os"
	"fmt"
	//"sort"
)

type word struct {
	spell string
	count, length int
}

func main() {
	var N, M int
	var s string
	var words map[string]int = make(map[string]int)
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &M)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &s)
		if len(s) >= M {
			words[s]++
		}
	}

	fmt.Println(words)

	var Word []word = make([]word, len(words))
	
	for i:=0;i<len(words);i++ { 
		for key, val := range words {
			fmt.Println(key, val)
			Word[i].spell = key
			Word[i].count = val
			Word[i].length = len(key)
		}
	}

	for i:=0;i<len(Word);i++ { 
		fmt.Println(Word[i].spell, Word[i].count, Word[i].length)
	}

	// sort.Slice(words, func(i, j int) bool {
		
	// })
}