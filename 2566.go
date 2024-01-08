package main

import (
	"fmt"
	//"bufio"
	//"os"
	"sort"
)

func main() {
	var matrix [9][9]int 
	var copy []int 
	//reader := bufio.NewReader(os.Stdin)

	for i:=0;i<9;i++ {
		for j:=0;j<9;j++ {
			fmt.Scan(&matrix[i][j])
			copy = append(copy, matrix[i][j])
		}
	}

	sort.Ints(copy)
	max := copy[len(copy)-1] 

	Loop1 :
	for i:=0;i<9;i++ {
		for j:=0;j<9;j++ {
			if matrix[i][j] == max {
				fmt.Println(max)
				fmt.Println(i+1, j+1)
				break Loop1
			}
		}
	}
}