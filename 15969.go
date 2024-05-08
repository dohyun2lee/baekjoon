package main 

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Fscanln(reader, &N)
	var score []int = make([]int, N)

	for i:=0;i<N;i++ {
		if i != N-1 {
			fmt.Fscan(reader, &score[i])
		} else {
			fmt.Fscanln(reader, &score[i])
		}
	}

	sort.Ints(score)

	fmt.Println(score[len(score)-1]-score[0])
}