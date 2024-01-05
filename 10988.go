package main 

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	var cnt int

	fmt.Scanln(&S)

	s := strings.Split(S, "") 

	for i:=0;i<len(s);i++ {
		if s[i] == s[len(s)-i-1] {
			cnt = 1
		} else {
			cnt = 0 
			break
		}
	}

	fmt.Println(cnt)
}