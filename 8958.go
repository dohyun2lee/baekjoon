package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var T int
	var S string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		var tmp, cnt int
		fmt.Fscanln(reader, &S)

		s := strings.Split(S, "")

		for j:=0;j<len(s);j++ {
			if s[j] == "O" {
				tmp++
				cnt += tmp
			} else {
				tmp = 0
			}
		}

		fmt.Fprintln(writer, cnt)
	}
}
