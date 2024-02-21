package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var cnt, ans int
	var Word []string
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()
	word := strings.Split(s, " ")

	for i := 0; i < len(word); i++ {
		Word = append(Word, strings.Split(word[i], "")...)
	}

	for i := 0; i < len(Word); i++ {
		if Word[i] == "U" && cnt == 0 {
			cnt++
		} else if Word[i] == "C" && cnt == 1 {
			cnt++
		} else if Word[i] == "P" && cnt == 2 {
			cnt++
		} else if Word[i] == "C" && cnt == 3 {
			cnt++
		}
		if cnt == 4 {
			ans = 1
			break
		}
	}

	if ans == 1 {
		fmt.Println("I love UCPC")
	} else {
		fmt.Println("I hate UCPC")
	}
}
