package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, M, pack, one, ans int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)
	var Pack []int = make([]int, M)
	var One []int = make([]int, M)

	for i := 0; i < M; i++ {
		fmt.Fscanln(reader, &pack, &one)

		Pack[i] = pack
		One[i] = one
	}

	sort.Ints(Pack)
	pack = Pack[0]
	sort.Ints(One)
	one = One[0]

	if N%6 == 0 {
		if N*one < N/6*pack {
			ans = N * one
		} else {
			ans = (N / 6) * pack
		}
	} else {
		if N%6*one < pack {
			if 6*one < pack {
				ans = N * one
			} else {
				ans = ((N / 6) * pack) + ((N % 6) * one)
			}
		} else {
			ans = (N/6 + 1) * pack
		}
	}

	fmt.Fprintln(writer, ans)
}
