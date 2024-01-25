package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
	"math"
)

func main() {
	var N, sum int
	max := 1
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var num []int = make([]int, N)
	var modechk map[int]int = make(map[int]int)
	var mode []int

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &num[i])
		sum += num[i]
		modechk[num[i]]++
	}

	sort.Ints(num)

	avg := math.Round((float64(sum)/float64(N)))

	if avg == -0 {
		fmt.Println(0)
	} else {
		fmt.Println(avg)
	}
	fmt.Println(num[N/2])
	
	for key, val := range modechk {
		if val > max {
			mode = []int{}
			mode = append(mode, key)
			max = val
		} else if val == max {
			mode = append(mode, key)
		}
	}
	sort.Ints(mode)
	if len(mode) > 1 {
		fmt.Println(mode[1])
	} else {
		fmt.Println(mode[0])
	}

	fmt.Println(num[N-1]-num[0])
}
