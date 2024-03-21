package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var S, ans string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &S)

	for i := 0; i < len(S); i++ {
		if S[i] < 97 {
			ans += string(S[i]+32)
		} else {
			ans += string(S[i]-32)
		}
	}

	fmt.Println(ans)
}
