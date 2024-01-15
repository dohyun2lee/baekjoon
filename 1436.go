package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var N, cnt int
	var i = 665
	fmt.Scanln(&N)

	for ;;i++ {
		a := strconv.Itoa(i)
		A := strings.Split(a, "")
		for j := 0; j <= len(A)-3; j++ {
			if A[j] == "6" && A[j+1] == "6" && A[j+2] == "6" {
				cnt++
				break
			}
		}
		if cnt == N {
			fmt.Println(i)
			break
		}
	}
}
