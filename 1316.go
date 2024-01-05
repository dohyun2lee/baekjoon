package main  

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	var N, cnt, ans int
	var S string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for a:=0;a<N;a++ {
		fmt.Fscanln(reader, &S)
		s := strings.Split(S, "")
		cnt = 0

		Loop1: 
		for i:=0;i<len(s);i++ {
			for j:=i+2;j<len(s);j++ {
				if s[i] == s[j] {
					if s[i] != s[i+1] {
						break Loop1
					}
				} 
			} 
			cnt++
		}

		if cnt == len(s) {
			ans++
		}
	}

	fmt.Println(ans)
}