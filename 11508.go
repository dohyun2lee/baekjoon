package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, sum, remainder int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var milk []int = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &milk[i])
	}

	sort.Ints(milk)

	remainder = N % 3

	if remainder == 1 {
		sum += milk[0]
		milk = milk[1:]
	} else if remainder == 2 {
		sum += milk[0] + milk[1]
		milk = milk[2:]
	}

	for i := 0; i < len(milk); i+=3 {
		sum += milk[i+1] + milk[i+2]
	}

	fmt.Println(sum)
}
