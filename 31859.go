package main 

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
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
	
	ans += strconv.Itoa(cnt+4)
	ans = strconv.Itoa(N+1906) + ans

	res = "smupc_"

	for i:=len(ans)-1;i>=0;i-- {
		res += string(ans[i])
	}

	fmt.Fprintf(writer, "%s", res)
}