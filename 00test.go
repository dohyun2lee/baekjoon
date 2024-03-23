package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Fscanln(reader, &N)

	switch N {
	case 1:
		fmt.Println("No")
	case 2:
	case 3:
		fmt.Println("hi")
	}
}
