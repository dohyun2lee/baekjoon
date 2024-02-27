package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var S string
	var cnt int
	reader := bufio.NewReader(os.Stdin)

	Loop1: for {
		cnt = 0
		fmt.Fscanln(reader, &S)

		s := strings.Split(S, "")

		for i:=0;i<len(s);i++ {
			if s[i] == "-" {
				break Loop1
			}

			if 
		}
	}
}
