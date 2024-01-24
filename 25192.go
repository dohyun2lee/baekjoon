	package main

	import (
		"bufio"
		"os"
		"fmt"
	)

	var person map[string]int = make(map[string]int)

	func main() {
		var N, cnt int
		var s string
		reader := bufio.NewReader(os.Stdin)

		fmt.Fscanln(reader, &N)
		
		for i:=0;i<N;i++ {
			fmt.Fscanln(reader, &s)

			if s == "ENTER" {
				reset()
			} else {
				if person[s] == 0 {
					cnt++
				}
				person[s] = 1
			}
		}

		fmt.Println(cnt)
	}

	func reset () {
		person = make(map[string]int)
	}