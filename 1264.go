package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	var cnt int
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		scanner.Scan()
		s := scanner.Text()
		cnt = 0

		if s == "#" {
			break
		} else {
			S := strings.Split(s, "")

			for i:=0;i<len(S);i++ {
				if S[i] == "a" || S[i] == "e" || S[i] == "i" || S[i] == "o" || S[i] == "u" || S[i] == "A" || S[i] == "E" || S[i] == "I" || S[i] == "O" || S[i] == "U" {
					cnt++
				}
			}
		}

		fmt.Fprintln(writer, cnt)
	}
}