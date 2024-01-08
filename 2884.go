package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var H, M int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &H, &M)

	if M >= 45 {
		fmt.Println(H, M-45)
	} else if H > 0 && M < 45{
		fmt.Println(H-1, M+15)
	} else {
		fmt.Println(23, M+15)
	}
}