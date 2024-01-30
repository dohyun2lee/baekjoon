package main

import "fmt"

func main(){
	var N, M int
	var matrix_one [101][101]int

	fmt.Scanln(&N, &M)

	for i:=0;i<N;i++ {
		for j:=0;j<M;j++ {
			fmt.Scan(&matrix_one[i][j])
		}
		fmt.Println()
	}

	fmt.Println(ma)
}