package main 

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var ans int
	var minus bool
	var s, old, new, sign string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &s)
	S := strings.Split(s, "")

	for i:=0;i<len(S);i++ {
		if S[i] == "+" {
			x, _ := strconv.Atoi(new)
			if minus {
				ans -= x
			} else {
				ans += x
			}
			old = new
			new = ""
		} else if S[i] == "-" {
			minus = true
			old = new
			new = ""
		} else {
			new += S[i]
		}
	}
}