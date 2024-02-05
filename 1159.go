package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var s string
	var N int
	var member map[string]int = make(map[string]int)
	var squad []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &s)

		member[string(s[0])]++
	}

	for key, val := range member {
		if val >= 5 {
			squad = append(squad, key)
		}
	}

	if len(squad) == 0 {
		fmt.Println("PREDAJA")
	} else {
		sort.Strings(squad)
		s = strings.Join(squad, "")
		fmt.Println(s)
	}
}
