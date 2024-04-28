package main 

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var sum, afterSum, mush int
	var mushroom []int
	reader := bufio.NewReader(os.Stdin)

	for i:=0;i<10;i++ {
		fmt.Fscanln(reader, &mush)
		mushroom = append(mushroom, mush)
	}

	for i:=0;i<10;i++ {
		afterSum = sum + mushroom[i]
		
		if afterSum <= 100 {
			sum = afterSum
			if sum == 100 {
				break
			}
		} else {
			if afterSum - 100 <= 100 - sum {
				sum = afterSum
				break
			} else {
				break
			}
		}
	}

	fmt.Println(sum)
}