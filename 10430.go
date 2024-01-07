package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var A, B, C int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A, &B, &C)

	fmt.Println((A+B)%C)
	fmt.Println(((A%C)+(B%C))%C)
	fmt.Println((A*B)%C)
	fmt.Println(((A%C)*(B%C))%C)
}