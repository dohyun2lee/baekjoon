package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N, Kim, Yim, cnt int
	reader := bufio.NewReader(os.Stdin)
	var people []int

	fmt.Fscanln(reader, &N, &Kim, &Yim)

	for i:=1;i<=N;i++ {
		people = append(people, i)
	}

	for {
		cnt++
		people, Kim, Yim = nxtRound(people, Kim, Yim)
		if Kim == Yim {
			break
		}
		people = makeSliceUnique(people)
	}

	fmt.Println(cnt)
}

func nxtRound(a []int, Kim, Yim int) ([]int, int, int) {
	var b []int
	for i:=0;i<len(a);i++ {
		if a[i]%2 != 0 {
			b = append(b, a[i]/2+1)
		} else {
			b = append(b, a[i]/2)
		}
	}
	Kim = b[Kim-1]
	Yim = b[Yim-1]
	
	return b, Kim, Yim
}

func makeSliceUnique(x []int) []int {
    keys := make(map[int]struct{}) 
    res := make([]int, 0) 
    for _, val := range x { 
        if _, ok := keys[val]; ok { 
            continue 
        } else { 
            keys[val] = struct{}{} 
            res = append(res, val) 
        } 
    }
    return res 
}