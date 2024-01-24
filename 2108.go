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
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var num []int = make([]int, N)
	var mode map[int]int = make(map[int]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &num[i])
		sum += num[i]
		mode[num[i]]++
	}

	sort.Ints(num)

	fmt.Println(math.Round((float64(sum)/float64(N))))
	fmt.Println(num[N/2])
	if modeChk(mode) == -1 {

	}
	fmt.Println(modeChk(mode))
	fmt.Println(num[(len(mode)-1)]-num[0])
}

func modeChk(mymap map[int]int) int {
	var ans_val, ans_key int
	var value []int
	var keys []int

	for key, val := range mymap {
		sort.Slice(mymap, func(key int, val int) bool {
			if mymap[i].val == mymap[j].val {
				return mymap[i].x < mymap[j].x
			}
			return mymap[i].y < mymap[j].y
		})	
	}
}