package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		S := scanner.Text()

		if S == "." {
			break
		}

		s := strings.Split(S, "")
		var chk int = 0
		var stack []string
		
		for j:=0;j<len(s);j++ {
			if s[j] == "(" {
				stack = append(stack, "(")
			} else if s[j] == ")" {
				if len(stack) != 0 && stack[len(stack)-1] != "(" {
					chk = -1
					break
				} else if len(stack) == 0 {
					chk = -1 
					break
				}
				stack = stack[:len(stack)-1]
			} else if s[j] == "[" {
				stack = append(stack, "[")
			} else if s[j] == "]" {
				if len(stack) != 0 && stack[len(stack)-1] != "[" {
					chk = -1
					break
				} else if len(stack) == 0 {
					chk = -1 
					break
				}
				stack = stack[:len(stack)-1]
			}
		}

		if len(stack) == 0 && chk == 0 {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}