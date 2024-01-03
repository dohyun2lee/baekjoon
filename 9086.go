package main 

import (
	"fmt"
	"strings"
)

func main() {
	var T int
	var S string

	fmt.Scanln(&T)

	for i:=0;i<T;i++ {
		fmt.Scanln(&S)
		s := strings.Split(S, "")
		ans := fmt.Sprintf("%s%s", s[0], s[len(s)-1])
		fmt.Println(ans)
	}
}