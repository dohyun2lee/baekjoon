package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var N, cnt int
	var x int = 1
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Fscanln(reader, &N)
	var line []int = make([]int, N)
	var side []int

	for i:=0;i<N;i++ {
		fmt.Fscanf(reader, "%d ", &line[i])
	}

	for {
		if len(line) > 0 && line[0] == x {
			x++
			cnt++
			line = line[1:]
		} else if len(side) > 0 && side[len(side)-1] == x {
			x++
			cnt++
			side = side[:len(side)-1]
		} else if len(line) == 0 && side[len(side)-1] != x {
			break
		} else {
			side = append(side, line[0])
			line = line[1:]
		} 

		if len(line) == 0 && len(side) == 0 {
			break
		}
	}

	if cnt == N {
		fmt.Println("Nice")
	} else {
		fmt.Println("Sad")
	}
}