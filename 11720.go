	package main

	import (
		"fmt"
		"strings"
		"strconv"
	)

	func main() {
		var N, sum int
		var X string

		fmt.Scanln(&N)
		fmt.Scanln(&X)

		x := strings.Split(X, "")

		for i:=0;i<N;i++ {
			a, _ := strconv.Atoi(x[i])
			sum = sum + a
		}

		fmt.Println(sum)
	}