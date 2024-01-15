package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a1, a0, c, n0 int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &a1, &a0)
	fmt.Fscanln(reader, &c)
	fmt.Fscanln(reader, &n0)

	if (a1*n0+a0 <= c*n0) && a1 <= c {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}

}
