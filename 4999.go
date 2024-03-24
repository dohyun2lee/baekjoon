package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var possible, required string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &possible)
	fmt.Fscanln(reader, &required)

	if len(possible) >= len(required) {
		fmt.Println("go")
	} else {
		fmt.Println("no")
	}
}
