package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var T int
	var S string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &T)

	for i:=0;i<T;i++ {
		fmt.Fscanln(reader, &S)
		s := strings.Split(S, "")
		var chkVPS int
		
		for j:=0;j<len(s);j++ {
			if s[j] == "(" {
				chkVPS++
			} else {
				chkVPS--
			}
			if chkVPS < 0 {
				chkVPS = -1
				break
			}
		}

		if chkVPS == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}