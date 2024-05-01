package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var A, I int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A, &I)

	fmt.Println((A*(I-1))+1)
}