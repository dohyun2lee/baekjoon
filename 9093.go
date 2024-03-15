package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 9093. 단어 뒤집기

func reverse(s string) string {

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()

	input := scanner.Text()
	t, _ := strconv.Atoi(input)

	for i := 0; i < t; i++ {
		scanner.Scan()
		sentence := scanner.Text()

		pieces := strings.Split(sentence, " ")
		for _, piece := range pieces {
			fmt.Fprint(writer, reverse(piece), " ")
		}

		fmt.Fprintln(writer)
	}
}