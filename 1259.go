package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		fmt.Fscanln(reader, &N)
		var ans bool = true

		if N == 0 {
			break
		}

		n := strconv.Itoa(N)

		num := strings.Split(n, "")
		tmp := len(num) - 1

		for i := 0; i < len(num)/2; i++ {
			if num[i] != num[tmp] {
				ans = false
				break
			}
			tmp--
		}

		if ans {
			fmt.Fprintln(writer, "yes")
		} else {
			fmt.Fprintln(writer, "no")
		}
	}
}
