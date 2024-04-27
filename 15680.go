package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	if N == 0 {
		fmt.Println("YONSEI")
	} else {
		fmt.Println("Leading the Way to the Future")
	}
}