package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var N, M int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &M)
	fmt.Fscanln(reader, &N)

	for i := M; i <= N; i++ {
		fmt.Println(int(math.Sqrt(float64(i))))
	}

}
