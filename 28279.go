package main

import (
	"bufio"
	"fmt"
	"os"
	"container/list"
)

func main() {
	var N int
	var cmd, x int
	listack := list.New()
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &cmd, &x)

		switch cmd {
		case 1:
			listack.PushFront(x)
		case 2:
			listack.PushBack(x)
		case 3:
			if listack.Len() == 0 {
				fmt.Fprintln(writer, -1)
			} else if listack.Len() == 1 {
				fmt.Fprintln(writer, listack.Front().Value)
				listack = list.New()
			} else {
				fmt.Fprintln(writer, listack.Front().Value)
				listack.Remove(listack.Front())
			}
		case 4:
			if listack.Len() == 0 {
				fmt.Fprintln(writer, -1)
			} else if listack.Len() == 1 {
				fmt.Fprintln(writer, listack.Back().Value)
				listack = list.New()
			} else {
				fmt.Fprintln(writer, listack.Back().Value)
				listack.Remove(listack.Back())
			}
		case 5:
			fmt.Fprintln(writer, listack.Len())
		case 6:
			if listack.Len() == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case 7:
			if listack.Len() == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, listack.Front().Value)
			}
		case 8:
			if listack.Len() == 0 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, listack.Back().Value)
			}
		}
	}
}
