package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, m, M, T, R, time, work, rate int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &m, &M, &T, &R)

	rate = m

	for {
		if m + T > M {
			time = -1
			break
		}
		
		if rate + T <= M {
			rate += T
			time++
			work++
		} else if rate + T > M && rate - R >= m {
			rate -= R
			time++
		} else if rate + T > M && rate - R < m {
			rate = m
			time++
		}

		if work == N {
			break
		}
	}

	fmt.Println(time)
}