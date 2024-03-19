package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i:=N;i>0;i-- {
		fmt.Fprintln(writer, i)
	}
}