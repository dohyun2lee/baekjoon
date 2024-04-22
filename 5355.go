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
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		scanner.Scan()
		x := strings.Split(scanner.Text(), " ")

		N, _ := strconv.ParseFloat(x[0], 32)
		for j := 1; j < len(x); j++ {

			if x[j] == "@" {
				N *= 3
			} else if x[j] == "%" {
				N += 5.0
			} else if x[j] == "#" {
				N -= 7.0
			}

		}

		fmt.Fprintf(writer, "%.2f\n", N)
	}
}
