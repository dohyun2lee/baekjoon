package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var a, min, d, e int
	reader := bufio.NewReader(os.Stdin)

	min = 2000

	for i:=0;i<3;i++ {
		fmt.Fscanln(reader, &a)

		if a < min {
			min = a
		}
	}

	fmt.Fscanln(reader, &d)
	fmt.Fscanln(reader, &e)

	if d > e {
		fmt.Println(min + e -50)
	} else {
		fmt.Println(min + d -50)
	}
}