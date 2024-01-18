package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N, M, cnt int
	var ans []string
	var s string
	var people map[string]int = make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &M)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &s)
		people[s]++
	}
	for i:=0;i<M;i++ {
		fmt.Fscanln(reader, &s)
		people[s]++
	}

	for key, val := range people {
		if val == 2 {
			ans = append(ans, key)
			cnt++
		}
	}

	sort.Strings(ans)
	
	fmt.Fprintln(writer, cnt)
	for i:=0;i<cnt;i++ {
		fmt.Fprintln(writer, ans[i])
	}
}