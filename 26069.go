package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var N int
	var a, b string
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &N)
	var person map[string]int = make(map[string]int)

	person["ChongChong"] = 1

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &a, &b)
		
		if person[a] == 0 && person[b] != 0 {
			person[a] = 1
		}else if person[b] == 0 && person[a] != 0 {
			person[b] = 1
		}
	}

	fmt.Println(chkRainbow(person))
}

func chkRainbow (mymap map[string]int) int {
	var cnt int

	for _, val := range mymap {
		if val == 1 {
			cnt++
		}
	}

	return cnt
}