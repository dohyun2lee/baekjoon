package main

import (
	// "bufio"
	// "os"
	"fmt"
)

func main() {
	var S = [3]int{1, 2, 3}
	var a = [3]int{1, 3, 3}

	if S == a {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
