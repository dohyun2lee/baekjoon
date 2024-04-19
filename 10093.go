package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var A, B int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &A, &B)
	// A, B 같을 때 고치기

	if B >= A {
		if B-A <= 1 {
			fmt.Fprintln(writer, 0)
		} else {
			fmt.Fprintln(writer, B-A-1)
	
			for i:=A+1;i<B;i++ {
				fmt.Fprintf(writer, "%d ", i)
			}
		}
	} else {
		if A-B <= 1 {
			fmt.Fprintln(writer, 0)
		} else {
			fmt.Fprintln(writer, A-B-1)
	
			for i:=B+1;i<A;i++ {
				fmt.Fprintf(writer, "%d ", i)
			}
		}
	}
}
