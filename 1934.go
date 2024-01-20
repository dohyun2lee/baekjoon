package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var T, A, B int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i:=0;i<T;i++ {
		fmt.Fscanln(reader, &A, &B)
		if A > B {
			for j:=1;;j++{
				if (A * j) % B == 0 {
					fmt.Fprintln(writer, A * j)
					break
				}
			}
		} else {
			for j:=1;;j++{
				if (B * j) % A == 0 {
					fmt.Fprintln(writer, B * j)
					break
				}
			}
		}
	}
}