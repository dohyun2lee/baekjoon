package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
)

type coordinate struct {
	x, y int
}

func main() {
	var N, x, y int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	var crd []coordinate = make([]coordinate, 0)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &x, &y)
		crd = append(crd, coordinate{x, y})
	}
	
	sort.Slice(crd, func(i int, j int) bool {
		if crd[i].x == crd[j].x {
			return crd[i].y < crd[j].y
		}
		return crd[i].x < crd[j].x
	})	

	for i:=0;i<N;i++ {
		fmt.Fprintln(writer, crd[i].x, crd[i].y)
	}
}


