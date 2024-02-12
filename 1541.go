package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var old int
	var minus bool
	var s, new string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &s)
	s = "+" + s
	S := strings.Split(s, "")

	for i := 0; i < len(S); i++ {
		if S[i] == "+" {
			if minus {
				for j := i + 1; j < len(S); j++ {
					if S[j] == "+" || S[j] == "-" {
						break
					}
					new += S[j]
				}
				x, _ := strconv.Atoi(new)
				old -= x
				new = ""
			} else {
				for j := i + 1; j < len(S); j++ {
					if S[j] == "+" || S[j] == "-" {
						break
					}
					new += S[j]
				}
				x, _ := strconv.Atoi(new)
				old += x
				new = ""
			}
		} else if S[i] == "-" {
			minus = true
			for j := i + 1; j < len(S); j++ {
				if S[j] == "+" || S[j] == "-" {
					break
				}
				new += S[j]
			}
			x, _ := strconv.Atoi(new)
			old -= x
			new = ""
		}
	}

	fmt.Println(old)
}
