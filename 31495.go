package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	S := scanner.Text()

	s := strings.Split(S, "")

	if s[0] == "\"" && s[len(s)-1] == "\"" && len(s) > 2{
		s = s[1:]
		s = s[:len(s)-1]
		fmt.Println(strings.Join(s, ""))
	} else {
		fmt.Println("CE")
	}
}
