package main

import (
	"fmt"
	"bufio"
	"os"
	//"sort"
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
	var crd []coordinate = make([]coordinate, N)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &x, &y)
		crd[i].x = x
		crd[i].y = y
	}
	
	Sorting(crd, N)

	for i:=0;i<N;i++ {
		fmt.Fprintln(writer, crd[i].x, crd[i].y)
	}
}


func Sorting(crd []coordinate, n int) []coordinate {
	for i:=0;i<n;i++ {
		for j:=i;j<n;j++ {
			if crd[i].x > crd[j].x {
				tmp := crd[i]
				crd[i] = crd[j]
				crd[j] = tmp
			} else if crd[i].x == crd[j].x && crd[i].y > crd[j].y {
				tmp := crd[i]
				crd[i] = crd[j]
				crd[j] = tmp
			}
		}
	}
	return crd
}