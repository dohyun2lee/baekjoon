package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	if a > 0 {
		if b > 0 {
			fmt.Println(1)
		} else {
			fmt.Println(4)
		}
	} else {
		if b > 0 {
			fmt.Println(2)
		} else {
			fmt.Println(3)
		}
	}
}