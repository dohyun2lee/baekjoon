package main

import (
	"bufio"
	"os"
	"fmt"
)

func main () {
	var N, F int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	fmt.Fscanln(reader, &F)	

	if (N-(N%100))%F == 0 {
		fmt.Println("00")
	} else {
		ans := (N-(N%100))/F*F+F
		if ans%100 < 10 {
			fmt.Printf("0%d", ans%100)
		} else {
			fmt.Printf("%d", ans%100)
		}
	}
}