package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var s string
	var ans int
	var words map[string]int = make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &s)

	S := strings.Split(s, "")

	for i:=0;i<len(S);i++ {
		for j:=i;j<len(S);j++ {
			words[s[i:j+1]]++
		}
	}

	for _, val := range words {
		if val > 0 {
			ans++
		}
	}

	fmt.Fprintln(writer, ans)
}