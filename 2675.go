package main

import (
	"fmt"
	"strings"
)

func main() {
	var T, R int 
	var S string

	fmt.Scanln(&T)

	for i:=0;i<T;i++ {
		fmt.Scanln(&R, &S)
		s := strings.Split(S, "")
		for j:=0;j<len(s);j++ {
			for k:=0;k<R;k++ {
				fmt.Print(s[j])
			}
		}
		fmt.Println()
	}
}