package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var A, B int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A)
	fmt.Fscanln(reader, &B)
	
	fmt.Println(A * (B%10))
	fmt.Println(A * ((B%100)/10))
	fmt.Println(A * (B/100))
	fmt.Println(A * B)
}