package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N, M int
	var site string
	var password map[string]string = make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)

	for i := 0; i < N; i++ {
		scanner.Scan()
		x := scanner.Text()
		X := strings.Split(x, " ")

		password[X[0]] = X[1]
	}

	for i := 0; i < M; i++ {
		fmt.Fscanln(reader, &site)
		fmt.Fprintln(writer, password[site])
	}
}
