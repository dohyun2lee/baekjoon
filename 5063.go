package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var T, r, e, c int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i:=0;i<T;i++ {
		fmt.Fscanln(reader, &r, &e, &c)

		if r > e - c {
			fmt.Fprintln(writer, "do not advertise")
		} else if r < e - c {
			fmt.Fprintln(writer, "advertise")
		} else {
			fmt.Fprintln(writer, "does not matter")	
		}
	}
}