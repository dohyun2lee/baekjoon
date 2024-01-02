package main 

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	var i int

	fmt.Scanln(&S)
	fmt.Scanln(&i)
	s := strings.Split(S, "")
	fmt.Print(s[i-1])
}