package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Fscanln(reader, &N)
	var board []int = make([]int, N+1)

	fmt.Println(board[0])
}
