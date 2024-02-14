package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var thr, N, Zero int
	var zero bool
	var numbers []string
	var ans string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for N>0{
		if N%10 == 0 {
			zero = true
		}
		numbers = append(numbers, strconv.Itoa(N%10))
		thr += N%10
		N /= 10
	}

	if thr%3 != 0 || zero != true {
		fmt.Println(-1)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(numbers)))
		for i:=0;i<len(numbers);i++ {
			ans += numbers[i]
		}
		res,_ := strconv.Atoi(ans)
		fmt.Println(res)
	}	
}
