package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &s)

	if s == "N" || s == "n" {
		fmt.Fprintln(writer, "Naver D2")
	} else {
		fmt.Fprintln(writer, "Naver Whale")
	}
}

