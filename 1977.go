package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var N, M, min, sum int
	reader := bufio.NewReader(os.Stdin)
	min = 10001

	fmt.Fscanln(reader, &M)
	fmt.Fscanln(reader, &N)

	for i := int(math.Sqrt(float64(M))); i <= int(math.Ceil(math.Sqrt(float64(N)))); i++ {
		if (i*i) >= M && (i*i) <= N {
			if i*i < min {
				min = i*i
			}
			sum += i * i
		}
	}

	if min > 10000 {
		fmt.Println(-1)
	} else {
		fmt.Println(sum)
		fmt.Println(min)
	}
}
