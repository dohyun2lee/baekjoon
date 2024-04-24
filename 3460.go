package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i:=0;i<T;i++ {
		fmt.Fscanln(reader, &n)
		
	}
}
