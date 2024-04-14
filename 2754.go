package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var sum float64
	var x string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &x)

	switch string(x[0]) {
	case "A":
		sum += 4
	case "B":
		sum += 3
	case "C":
		sum += 2
	case "D":
		sum += 1
	case "F":
		sum = 0
	}

	if sum == 0 {
		fmt.Println("0.0")
	} else {
		switch string(x[1]) {
		case "+":
			sum += 0.3
		case "-":
			sum -= 0.3
		}
		fmt.Printf("%.1f", sum)
	}
}
