package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var H, M, T int
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Fscanln(reader, &H, &M)
	fmt.Fscanln(reader, &T)

	if M+T >= 60 && H+((M+T)/60) < 24 {
		fmt.Println(H+((M+T)/60), (M+T)%60)
	} else if M+T < 60 {
		fmt.Println(H, M+T)
	} else {
		fmt.Println((((M+T)/60)+H)%24, (M+T)%60)
	}

}