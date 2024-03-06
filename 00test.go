package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var K, N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &K, &N)
	var lan []int = make([]int, N)

	for i := 0; i < K; i++ {
		fmt.Fscanln(reader, &lan[i])
	}

	fmt.Println(lan)
}
