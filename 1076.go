package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var color string
	var sum int
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 3; i++ {
		fmt.Fscanln(reader, &color)
		switch color {
		case "black":
			if i == 2 {
				sum *= 1
			} else {
				sum += 0
			}
		case "brown":
			if i == 2 {
				sum *= 10
			} else {
				sum += 1
			}
		case "red":
			if i == 2 {
				sum *= 100
			} else {
				sum += 2
			}
		case "orange":
			if i == 2 {
				sum *= 1000
			} else {
				sum += 3
			}
		case "yellow":
			if i == 2 {
				sum *= 10000
			} else {
				sum += 4
			}
		case "green":
			if i == 2 {
				sum *= 100000
			} else {
				sum += 5
			}
		case "blue":
			if i == 2 {
				sum *= 1000000
			} else {
				sum += 6
			}
		case "violet":
			if i == 2 {
				sum *= 10000000
			} else {
				sum += 7
			}
		case "grey":
			if i == 2 {
				sum *= 100000000
			} else {
				sum += 8
			}
		case "white":
			if i == 2 {
				sum *= 1000000000
			} else {
				sum += 9
			}
		}

		if i == 0 {
			sum *= 10
		}
	}

	fmt.Println(sum)
}
