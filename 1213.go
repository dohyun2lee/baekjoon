package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type word struct {
	spell string
	count int
}

func main() {
	var S, ans, chk string
	var Ans []string
	var cnt int
	var wordTable []word
	var letter map[string]int = make(map[string]int)
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &S)
	s := strings.Split(S, "")

	for i := 0; i < len(s); i++ {
		letter[s[i]]++
	}

	for k, v := range letter {
		if letter[k]%2 == 1 {
			cnt++
		}
		if cnt > 1 {
			break
		}
		wordTable = append(wordTable, word{k, v})
	}

	fmt.Println(wordTable)

	if cnt > 1 {
		fmt.Println("I'm Sorry Hansoo")
	} else {
		sort.Slice(wordTable, func(i, j int) bool {
			if wordTable[i].count/2 == wordTable[j].count/2 {
				return wordTable[i].spell < wordTable[j].spell
			}
			return wordTable[i].count/2 > wordTable[j].count/2
		})

		fmt.Println(wordTable)

		for i := 0; i < len(wordTable); i++ {
			if wordTable[i].count%2 == 0 {
				for j := 0; j < wordTable[i].count/2; j++ {
					Ans = append(Ans, wordTable[i].spell)
				}
			} else {
				for j := 0; j < wordTable[i].count/2; j++ {
					Ans = append(Ans, wordTable[i].spell)
				}
				chk = wordTable[i].spell
			}
		}

		fmt.Println(Ans)

		for i := 0; i < len(Ans); i++ {
			ans += Ans[i]
		}
		ans += chk
		for i := len(Ans) - 1; i >= 0; i-- {
			ans += Ans[i]
		}

		fmt.Println(ans)
	}
}
