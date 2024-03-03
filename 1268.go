package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, max, ans int
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var class [][]int = make([][]int, N)
	var stu []int = make([]int, N)
	var student []int = make([]int, 5)

	for i := 0; i < N; i++ {
		for j := 0; j < 5; j++ {
			if j != 4 {
				fmt.Fscan(reader, &student[j])
			} else {
				fmt.Fscanln(reader, &student[j])
			}
		}
		class[i] = append(class[i], student...)
	}

	for i := 0; i < N; i++ {
		var tmp []int
		for j := 0; j < 5; j++ {
			for k := 0; k < N; k++ {
				if i != k && class[i][j] == class[k][j] {
					tmp = append(tmp, k)
				}
			}
		}
		stu[i] = len(makeSliceUnique(tmp))
	}

	for i:=0;i<N;i++ {
		if max < stu[i] {
			max = stu[i]
			ans = i
		}
	}

	fmt.Println(ans+1)
}

func makeSliceUnique(s []int) []int {
    keys := make(map[int]struct{}) 
    res := make([]int, 0) 
    for _, val := range s { 
        if _, ok := keys[val]; ok { 
            continue 
        } else { 
            keys[val] = struct{}{} 
            res = append(res, val) 
        } 
    }
    return res 
}