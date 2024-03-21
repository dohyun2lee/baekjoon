package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N string
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Fscanln(reader, &N)

	fmt.Println(N)
}
