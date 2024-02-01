package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	var N int
	var s, ans string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var cmd [][]string = make([][]string, N)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &s)
		cmd[i] = strings.Split(s, "")
	}

	S := cmd[0]

	for i:=1;i<N;i++ {
		for j:=0;j<len(cmd[i]);j++ {
			if S[j] != cmd[i][j] {
				S[j] = "?"
			}
		}
	}
	
	for i:=0;i<len(S);i++ {
		ans += S[i]
	}

	fmt.Println(ans)
}