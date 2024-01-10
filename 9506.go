package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)

	for {
		var sum int
		var divior []int
		fmt.Fscanln(reader, &n)

		if n == -1 {
			break
		}

		for i:=1;i<n;i++{
			if n%i == 0 {
				sum += i
				divior = append(divior, i)
			}
		}
		if n == sum {
			fmt.Printf("%d = ", n)
			for j:=0;j<len(divior);j++ {
				fmt.Printf("%d", divior[j])
				if j != len(divior)-1 {
					fmt.Print(" + ")
				}
			}
			fmt.Println()
		} else {
			fmt.Printf("%d is NOT perfect.\n", n)
		}
	}
}