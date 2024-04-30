package main 

import (
	"bufio"
	"os"
	"fmt"
	//"math"
)

func main() {
	var H, W, M, N int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &H, &W, &N, &M)

	x := ((H-1)/(N+1)) + 1
	y := ((W-1)/(M+1)) + 1

	fmt.Println(x * y)
}