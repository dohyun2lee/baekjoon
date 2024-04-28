package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var sum int
	var S string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &S)

	for i:=0;i<len(S);i++ {
		if i == 0 {
			sum += 10
		} else {
			if S[i] == S[i-1] {
				sum += 5
			} else {
				sum += 10
			}
		}
	}

	fmt.Println(sum)
}