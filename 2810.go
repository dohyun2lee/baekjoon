package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var S string
	var N int
	var cnt int = 1
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	fmt.Fscanln(reader, &S)

	seat := strings.Split(S, "")

	for i := 0; i < len(seat); i++ {
		if seat[i] == "S" {
			cnt++
		} else {
			cnt++
			i++
		}
	}

	if cnt > N {
		fmt.Println(N)
	} else {
		fmt.Println(cnt)
	}
}
