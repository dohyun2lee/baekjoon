package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N, ans int
	var s1, s2 string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	fmt.Fscanln(reader, &s1)
	fmt.Fscanln(reader, &s2)

	if s1 == s2 {
		ans = 0
	} else {
		start := strings.Split(s1, "")

		for {
			if 
			ans++
			fmt.Println("start:",strings.Join(start, ""))
			fmt.Println("s2:", s2)
			if strings.Join(start, "") == s2 {
				break
			}
		}
	}

	fmt.Println(ans)
}
