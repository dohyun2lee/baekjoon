package main 

import (
	"bufio"
	"os"
	"fmt"
	"math"
)

func main() {
	var N, dalgu, poniex, diff, gcd int
	//var S string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var score []string = make([]string, N)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &score[i])	
	}

	for i:=0;i<N;i++ {
		if score[i] == "D" {
			dalgu++
		} else {
			poniex++
		}
		
		diff = int(math.Abs(float64(dalgu)-float64(poniex)))
		
		if diff >= 2 {
			break
		}
	}
	
	if dalgu == 0 || poniex == 0 {
		gcd = 1
	} else {
		gcd = getGCD(dalgu, poniex)
	}

	fmt.Printf("%d:%d\n", dalgu/gcd, poniex/gcd)
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