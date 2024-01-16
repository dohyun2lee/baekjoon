package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

func main() {
	var N int
	var slice []int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i:=0;i<N;i++ {
		var a int
		fmt.Fscanln(reader, &a)
		slice = append(slice, a)
	}

	sort.Ints(slice)

	for i:=0;i<N;i++ {
		fmt.Fprintln(writer, slice[i])
	}
}