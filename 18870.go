package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var X []int = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscanf(reader, "%d ", &X[i])
	}

	x := uniqueSlice(X)
	sort.Ints(x)
	var mymap map[int]int = make(map[int]int, len(x))
	for i:=0;i<len(x);i++ {
		mymap[x[i]] = i
	}

	for i:=0;i<N;i++ {
		fmt.Fprintf(writer, "%d ", mymap[X[i]])
	}
}

func uniqueSlice(s []int) []int {
    var res = map[int]int{}
	var ans []int

	for i:=0;i<len(s);i++ {
		if res[s[i]] == 0 {
			res[s[i]]++
			ans = append(ans, s[i])
		}
	}

	return ans
}