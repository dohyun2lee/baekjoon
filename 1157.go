package main

import (
	"fmt"
	"strings"
	"sort"
	"bufio"
	"os"
)

func main() {
	var S, max_alp string
	var max_num int
	var cnt int = 1

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &S)

	LS := strings.ToLower(S)
	s:= strings.Split(LS, "")
	sort.Strings(s)

	if len(s) == 1 {
		fmt.Println(strings.ToUpper(s[0]))
	}

	for i:=0;i<len(s)-1;i++ {	
		if s[i] == s[i+1] {
			cnt++
			if cnt > max_num {
				max_alp = s[i]
				max_num = cnt
			} else if cnt == max_num {
				max_alp = "?"
			}
		} else {
			if cnt > max_num {
				max_num = cnt
				max_alp = s[i]
			}
			cnt = 1
		}
	}

	if max_num == 1 {
		fmt.Println("?")
	} else {
		fmt.Println(strings.ToUpper(max_alp))
	}

}