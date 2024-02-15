package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var S, sign string
	var zero, one int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &S)

	s := strings.Split(S, "")

	sign = s[0]

	if s[len(s)-1] == "0" {
		s = append(s, "1")
	} else {
		s = append(s, "0")
	}

	for i := 0; i < len(s); i++ {
		if s[i] != sign {
			if sign == "0" {
				zero++
				sign = "1"
			} else {
				one++
				sign = "0"
			}
		}
	}

	if zero < one {
		fmt.Println(zero)
	} else {
		fmt.Println(one)
	}
}
