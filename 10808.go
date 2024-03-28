package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var S string
	var alpa map[string]int = make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &S)
	SS := strings.Split(S, "")

	for i := 0; i < len(SS); i++ {
		alpa[SS[i]]++
	}

	for i := 97; i < 123; i++ {
		fmt.Fprintf(writer, "%d ", alpa[string(i)])
	}
}
