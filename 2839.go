package main

import (
	"fmt"
)

func main() {
	var N, cnt int

	fmt.Scanln(&N)

	for {
		if (N%5)%3 == 0 {
			cnt = cnt + (N / 5) + ((N % 5) / 3)
			break
		} else {
			if N>3 {
				cnt++ 
				N -= 3
			} else {
				cnt = -1
				break
			}
		}
	}
	
	fmt.Println(cnt)
}
