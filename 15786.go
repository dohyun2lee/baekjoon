package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, M int
	var key, password string
	var ans bool
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)
	fmt.Fscanln(reader, &key)

	for i := 0; i < M; i++ {
		var val = 0
		ans = false

		fmt.Fscanln(reader, &password)

		for j := 0; j < len(password); j++ {
			if key[val] == password[j] {
				val++
			}

			if val == len(key) {
				ans = true
				break
			}
		}

		fmt.Fprintln(writer, ans)
	}
}
