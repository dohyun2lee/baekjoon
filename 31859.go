package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, cnt int
	var S, res, ans string
	var ap map[string]bool = make(map[string]bool)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &S)

	for i:=0;i<len(S);i++ {
		if ap[string(S[i])] == false {
			ans += string(S[i])
			ap[string(S[i])] = true
		} else {
			cnt++
		}
	}

	ans += string(cnt+4)
	ans = string(N+1906) + ans

	for i:=len(S)-1;i>=0;i-- {
		res += string(ans[i])
	}

	fmt.Fprintf(writer, "smupc_%s", res)
}