package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var L, n, ans int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &L)
	var S []int = make([]int, L)

	for i := 0; i < L; i++ {
		if i == L-1 {
			fmt.Fscanln(reader, &S[i])
		} else {
			fmt.Fscan(reader, &S[i])
		}
	}

	fmt.Fscanln(reader, &n)

	sort.Ints(S)

	if S[0] == n {
		ans = 0
	} else if S[0] > n {
		ans = (n - 1) + (S[0] - n - 1) + ((n - 1) * (S[0] - n - 1))
	} else {
		for i := 1; i < L; i++ {
			if S[i] == n {
				ans = 0
				break
			}
			if S[i] > n {
				ans = (n - S[i-1] - 1) + (S[i] - n - 1) + ((n - S[i-1] - 1) * (S[i] - n - 1))
				break
			}
		}
	}

	fmt.Println(ans)
}
