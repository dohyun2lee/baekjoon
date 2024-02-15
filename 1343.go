package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var S, ans string
	var num []int
	var cnt int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &S)

	s := strings.Split(S, "")
	s = append(s, ".")

	for i := 0; i < len(s); i++ {
		if s[i] == "X" {
			cnt++
		} else {
			if cnt > 0 {
				num = append(num, cnt)
			}
			cnt = 0
			num = append(num, cnt)
		}
	}

	for i := 0; i < len(num)-1; i++ {
		if num[i]%2 != 0 {
			ans = "-1"
			break
		} else if num[i] == 0 {
			ans += "."
		} else {
			var chk int = num[i]
			for chk > 0 {
				if chk >= 4 {
					ans += "AAAA"
					chk -= 4
				} else {
					ans += "BB"
					chk -= 2
				}
			}
		}
	}

	fmt.Println(ans)
}
