package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

func main() {
	var N int
	var name, info string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var enter = map[string]int{}
	var res []string

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &name, &info)
		if info == "enter" {
			enter[name]++
		} else {
			enter[name]--
		}
	}

	for k, _ := range enter {
		if enter[k] > 0 {
			res = append(res, k)
		}
	}	

	sort.Strings(res)

	for i:=len(res)-1;i>=0;i-- {
		fmt.Println(res[i])
	}
}
