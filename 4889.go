package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var S string
	var cnt, plus, x int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

Loop1:
	for {
		x++
		cnt = 0
		plus = 0
		fmt.Fscanln(reader, &S)

		s := strings.Split(S, "")

		for i := 0; i < len(s); i++ {
			if s[i] == "-" {
				break Loop1
			}

			if s[i] == "{" {
				plus++
			} else {
				if plus > 0 {
					plus--
				} else {
					cnt++
					plus++
				}
			}
		}

		cnt += plus / 2

		fmt.Fprintf(writer, "%d. %d\n", x, cnt)
	}
}
