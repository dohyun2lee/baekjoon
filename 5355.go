package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var T int
	var N float64
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		scanner.Scan()
		x := strings.Split(scanner.Text(), " ")

		fmt.Println(x)

		for j := 0; j < len(x); j++ {
			if j == 0 {
				N, _ = strconv.ParseFloat(x[j], 32)
			} else {
				if x[j] == "@" {
					N *= 3
				} else if x[j] == "%" {
					N += 5
				} else if x[j] == "#" {
					N -= 7
				}
			}
		}

		fmt.Fprintf(writer, "%.2f\n", N)
	}
}
