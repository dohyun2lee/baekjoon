package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	if N == 1 {
		fmt.Println(1)
	} else if N == 2 {
		fmt.Println(2)
	} else {
		var card []int = make([]int, N)
		
		for i:=0;i<N;i++ {
			card[i] = i+1
		}
	
		for {
			x := card[1]
			if len(card) == 2 {
				fmt.Println(card[1])
				break
			}
			card = card[2:]
			card = append(card, x)
		}
	}
}
