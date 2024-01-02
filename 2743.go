package main 

import (
	"fmt"
	"strings"
)

func main() {
	var S string

	fmt.Scanln(&S)
	s := strings.Split(S, "")
	fmt.Print(len(s))
}