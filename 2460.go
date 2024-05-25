package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var max, out, in, person int
	reader := bufio.NewReader(os.Stdin)

	for i:=0;i<10;i++ {
		fmt.Fscanln(reader, &out, &in)

		person -= out
		person += in

		if person > max {
			max = person
		}
	}

	fmt.Println(max)
}