package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var sum, diff, mush int
	reader := bufio.NewReader(os.Stdin)

	for i:=0;i<10;i++ {
		fmt.Fscanln(reader, &mush)
		
		if sum + mush <= 100 {
			sum += mush
			diff = 100 - sum
		} else {
			if sum + mush - 100 <= diff {
				sum += mush
			}
		}
	}

	fmt.Println(sum)
}