package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

func main() {
	var N, M, min, Max int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &M)
	var A []int = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscan(reader, &A[i])
	}

	sort.Ints(A)
	min = M

	for i:=0;i<N;i++ {
		for j:=i+1;j<N;j++ {
			for k:=j+1;k<N;k++ {
				if M - (A[i] + A[j] + A[k]) < min && M - (A[i] + A[j] + A[k]) >= 0{
					min = M - (A[i] + A[j] + A[k])
					Max = A[i] + A[j] + A[k]
				} 
				if M - (A[i] + A[j] + A[k]) < 0 {
					break
				}
			}
		}
	}

	fmt.Println(Max)
}