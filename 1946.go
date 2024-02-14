package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type score struct {
	document  int
	interview int
}

func main() {
	var T, N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		var cnt int
		fmt.Fscanln(reader, &N)
		var person []score = make([]score, N)
		for j := 0; j < N; j++ {
			fmt.Fscanln(reader, &person[j].document, &person[j].interview)
		}

		sort.Slice(person, func(i, j int) bool {
			if person[i].document == person[j].document {
				return person[i].interview < person[j].interview
			}
			return person[i].document < person[j].document
		})

		for j := 0; j < N; j++ {
			var ok_cnt bool = true
			for k := j; k >= 0; k-- {
				if person[j].interview > person[k].interview {
					ok_cnt = false
					break
				}
			}
			if ok_cnt {
				cnt++
			}
		}

		fmt.Fprintln(writer, cnt)
	}
}

