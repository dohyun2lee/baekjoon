package main

import (
	"bufio"
	"fmt"
	"os"
)

var cnt int

func main() {
	var N, K int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N, &K)
	var arr []int = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscanf(reader, "%d ", &arr[i])
	}


}

func merge_sort(A []int) { 
	p = 
    if p < r {
        q = (p + r) / 2   
        merge_sort(A, p, q)  
        merge_sort(A, q + 1, r)
        merge(A, p, q, r)
    }
}

func merge(A []int, p, q, r int) {
    i := p
	j := q + 1 
	t := 1;
    for i <= q && j <= r {
        if A[i] <= A[j] {
			tmp[t+1] = A[i+1]
		} else {
			tmp[t+1] = A[j+1]
		}
    }
	for i <= q {
		tmp[t+1] = A[i+1]
	}
for j <= r {
	tmp[t+1] = A[j+1]
}
    i = p
	t = 1

	for i <= r {
		A[i+1] = tmp[t+1]
	}

}