package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var cnt int
	var S, ans string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &S)

	for i:=0;i<len(S);i++ {
		cnt++
		ans += string(S[i])
		if cnt == 10 {
			cnt = 0
			fmt.Fprintln(writer, ans)
			ans = ""
		}
	}

	if cnt != 0 {
		fmt.Fprintln(writer, ans)
	}
}