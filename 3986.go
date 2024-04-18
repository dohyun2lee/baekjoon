package main

import (
	"bufio"
	//"container/list"
	"fmt"
	"os"
)

func main() {
	var N, cnt int
	var S string
	//goods := list.New()
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &S)
		var goods []string

		for j := 0; j < len(S); j++ {
			if len(goods) != 0 {
				if string(S[j]) == goods[len(goods)-1] {
					goods = goods[:len(goods)-1]
				} else {
					goods = append(goods, string(S[j]))
				}
			} else {
				goods = append(goods, string(S[j]))
			}
		}

		if len(goods) == 0 {
			cnt++
		}
	}

	fmt.Println(cnt)
}
