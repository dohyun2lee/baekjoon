package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var A, B, D int
	var aScore, bScore []int = make([]int, 10), make([]int, 10)
	reader := bufio.NewReader(os.Stdin)

	for i:=0;i<10;i++ {
		if i != 9 {
			fmt.Fscan(reader, &aScore[i])
		} else {
			fmt.Fscanln(reader, &aScore[i])
		}
	}
	for i:=0;i<10;i++ {
		if i != 9 {
			fmt.Fscan(reader, &bScore[i])
		} else {
			fmt.Fscanln(reader, &bScore[i])
		}
	}
	for i:=0;i<10;i++ {
		if aScore[i] > bScore[i] {
			A++
		} else if bScore[i] > aScore[i] {
			B++ 
		} else {
			D++
		}
	}

	if A > B {
		fmt.Println("A")
	} else if B > A {
		fmt.Println("B")
	} else {
		fmt.Println("D")
	}
}