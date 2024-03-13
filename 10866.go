package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	var N, x int
	var cmd string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	l := list.New()

	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &cmd, &x)

		switch cmd {
		case "push_front":
			l.PushFront(x)
		case "push_back":
			l.PushBack(x)
		case "pop_front":
			a := l.Front()
			if a == nil {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, a.Value)
				l.Remove(l.Front())
			}
		case "pop_back":
			a := l.Back()
			if a == nil {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, a.Value)
				l.Remove(l.Back())
			}
		case "size":
			fmt.Fprintln(writer, l.Len())
		case "empty":
			if l.Len() == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case "front":
			a := l.Front()
			if a == nil {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, a.Value)
			}
		case "back":
			a := l.Back()
			if a == nil {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, a.Value)
			}
		}
	}
}
