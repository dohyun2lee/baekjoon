package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type word struct {
	s string
	len int
}

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var words []string

	for i:=0;i<N;i++ {
		var s string
		fmt.Fscanln(reader, &s)
		words = append(words, s)
	}

	uniqueWords := uniqueSlice(words)
	var slice []word = make([]word, len(uniqueWords))

	for i:=0;i<len(uniqueWords);i++ {	// 	
		slice[i].s = uniqueWords[i]
		slice[i].len = len(uniqueWords[i])
	}

	sort.Slice(slice, func(i, j int) bool {
		if slice[i].len == slice[j].len {
			return slice[i].s < slice[j].s
		}
		return slice[i].len < slice[j].len
	})

	for i:=0;i<len(slice);i++ {
		fmt.Fprintln(writer, slice[i].s)
	}
}

func uniqueSlice(s []string) []string {
    var res = map[string]int{}
	var ans []string

	for i:=0;i<len(s);i++ {
		if res[s[i]] == 0 {
			res[s[i]]++
			ans = append(ans, s[i])
		}
	}

	return ans
}