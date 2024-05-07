package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var Monster, DI, DIP float64
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &Monster, &DI)

	DIP = Monster - (Monster*DI/100)

	if DIP >= 100 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}