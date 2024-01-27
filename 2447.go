package main

import (
	"bufio"
	"os"
	"fmt"
	"math"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	n := math.Log(3) 
}