package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var A, B string
	var tmp []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A)
	fmt.Fscanln(reader, &B)

	a := strings.Split(A, "")
	b := strings.Split(B, "")

	for {
		if b[len(b)-1] == "A" {
			b = b[:len(b)-1]
		} else {

		}
	}
}
