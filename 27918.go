package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, dalgu, poniex int
	var S string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &S)

		if S == "D" {
			dalgu++
		} else {
			poniex++
		}
	}

	gcd := getGCD(dalgu, poniex)

	fmt.Printf("%d:%d", dalgu/gcd, poniex/gcd)
}

func getGCD(first, second int) (gcd int) {
	if first < second {
		second, first = first, second
	}

	for second != 0 {
		first, second = second, first%second
	}
	return first
}